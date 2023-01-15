package kafka

import (
	"account-consumer-service/internal/pkg/utils"
	"time"

	"github.com/Shopify/sarama"
)

type IProducer struct {
	syncProducer sarama.SyncProducer
	schema       *SchemaRegistry
}

func (kc *KafkaClient) NewProducer() *IProducer {
	producer, err := sarama.NewSyncProducerFromClient(kc.Client)
	if err != nil {
		utils.Logger.Fatal("Error during kafka producer. Details: %v", err)
		panic(producer)
	}
	return &IProducer{producer, kc.SchemaRegistry}
}

func (ip *IProducer) Send(msg interface{}, topic, subject string) error {
	msgEncoder, err := ip.schema.Encode(msg, subject)
	if err != nil {
		utils.Logger.Error("Error send msg: %v", err)
		return err
	}

	m := sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(time.Now().String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}
	ip.syncProducer.SendMessage(&m)

	return nil
}
