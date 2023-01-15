package kafka

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/VanessaVallarini/account-toolkit/avros"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const (
	topic_account_createorupdate_dlq = "account_createorupdate_dlq"
	topic_account_delete_dlq         = "account_delete_dlq"
	topic_account_createorupdate     = "account_createorupdate"
	topic_account_delete             = "account_delete"
)

type Consumer struct {
	ready          chan bool
	dlqTopic       []string
	consumerTopic  []string
	sr             *SchemaRegistry
	producer       *IProducer
	accountService *services.AccountService
}

func NewConsumer(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *KafkaClient, accountService *services.AccountService) error {

	kafkaProducer := kafkaClient.NewProducer()

	consumer := Consumer{
		ready:          make(chan bool),
		dlqTopic:       cfg.DlqTopic,
		consumerTopic:  cfg.ConsumerTopic,
		sr:             kafkaClient.SchemaRegistry,
		producer:       kafkaProducer,
		accountService: accountService,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ctx := context.Background()
			propagators := propagation.TraceContext{}
			handler := otelsarama.WrapConsumerGroupHandler(&consumer, otelsarama.WithPropagators(propagators))
			if err := kafkaClient.GroupClient.Consume(ctx, cfg.ConsumerTopic, handler); err != nil {
				utils.Logger.Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				err := kafkaClient.GroupClient.Close()
				if err != nil {
					utils.Logger.Errorf("Error from consumer: %v", err)
				}

				utils.Logger.Info("consume closed, consuming again")
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	utils.Logger.Info("Sarama consumer up and running!...")

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
		utils.Logger.Fatal("Error closing groupClient: %v", err)
		panic(kafkaClient.GroupClient.Close())
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
			ctx := context.Background()
			if err := consumer.processMessage(ctx, message); err != nil {
				consumer.sendToDlq(ctx, consumer.dlqTopic, message)
			}
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func (consumer *Consumer) sendToDlq(ctx context.Context, dlqTopic []string, message *sarama.ConsumerMessage) {
	topic := consumer.getTopicDlq(message)

	subject := consumer.getSubject(topic)

	_, span := otel.GetTracerProvider().Tracer("consumer").Start(ctx, "sendToDlq")
	defer span.End()
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Value),
		Timestamp: time.Now(),
	}
	for _, header := range message.Headers {
		msg.Headers = append(msg.Headers, *header)
	}

	consumer.producer.Send(msg, topic, subject)
}

func (consumer *Consumer) getTopicDlq(message *sarama.ConsumerMessage) string {
	switch message.Topic {
	case topic_account_createorupdate:
		return topic_account_createorupdate_dlq
	case topic_account_delete:
		return topic_account_delete_dlq
	}

	utils.Logger.Errorf("DLQ topic not found. Topic message: %v", message.Topic)
	return ""
}

func (consumer *Consumer) getSubject(dlqTopic string) string {
	switch dlqTopic {
	case topic_account_createorupdate_dlq:
		return avros.AccountCreateOrUpdateSubject
	case topic_account_delete_dlq:
		return avros.AccountDeleteSubject
	}

	utils.Logger.Errorf("DLQ topic invalid. Topic: %v", dlqTopic)
	return ""
}
