package kafka

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

type Consumer struct {
	ready                  chan bool
	dlqTopic               string
	consumerTopic          string
	sr                     *SchemaRegistry
	producer               sarama.SyncProducer //usar no futuro para enviar para DLQ
	accountServiceConsumer *services.AccountService
}

func NewConsumer(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *KafkaClient, asc *services.AccountService) error {
	consumer := Consumer{
		sr:                     kafkaClient.SchemaRegistry,
		ready:                  make(chan bool),
		dlqTopic:               cfg.DlqTopic,
		consumerTopic:          cfg.ConsumerTopic,
		accountServiceConsumer: asc,
	}

	wg := &sync.WaitGroup{} //tratar erros em go rotinas de forma concorrente
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ctx := context.Background()
			propagators := propagation.TraceContext{}
			handler := otelsarama.WrapConsumerGroupHandler(&consumer, otelsarama.WithPropagators(propagators))
			if err := kafkaClient.GroupClient.Consume(ctx, []string{cfg.ConsumerTopic}, handler); err != nil {
				zap.S().Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				err := kafkaClient.GroupClient.Close()
				if err != nil {
					zap.S().Fatalf("Error from consumer: %v", err)
				}

				zap.S().Info("consume closed, consuming again")
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	zap.S().Info("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		log.Println("terminating: via signal")
	}

	wg.Wait()
	if err := kafkaClient.GroupClient.Close(); err != nil {
		zap.S().Panicf("Error closing groupClient: %v", err)
	}
	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			var ac models.AccountCreateEvent
			if err := consumer.sr.Decode(message.Value, &ac, models.AccountCreateSubject); err != nil {
				utils.Logger.Warn("error during decode message consumer kafka")
			}
			ctx := context.Background()
			consumer.accountServiceConsumer.CreateAccount(ctx, ac)
			fmt.Println(ac)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func (consumer *Consumer) sendToDlq(ctx context.Context, dlqTopic string, message *sarama.ConsumerMessage) {
	ctx, span := otel.GetTracerProvider().Tracer("consumer").Start(ctx, "sendToDlq")
	defer span.End()
	msg := &sarama.ProducerMessage{
		Topic:     dlqTopic,
		Key:       sarama.ByteEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Value),
		Timestamp: time.Now(),
	}
	for _, header := range message.Headers {
		msg.Headers = append(msg.Headers, *header)
	}

	partition, offset, err := consumer.producer.SendMessage(msg)
	if err != nil {
		zap.S().Error(err)
		span.SetStatus(codes.Error, err.Error())
		// change to retry queues instead of recursive approach
		consumer.sendToDlq(ctx, dlqTopic, message)
	}
	span.SetAttributes(attribute.String("topic", dlqTopic))
	span.SetAttributes(attribute.Int("partition", int(partition)))
	span.SetAttributes(attribute.Int64("offset", offset))
	zap.S().Infof("Message sent to dlq: topic = %s, partition = %v, offset = %v", dlqTopic, partition, offset)
}
