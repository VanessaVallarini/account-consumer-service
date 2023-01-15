package kafka

import (
	"account-consumer-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKafkaClient(t *testing.T) {
	configKafka := models.KafkaConfig{
		ClientId:               "account-consumer-service",
		Hosts:                  []string{"localhost:9092"},
		SchemaRegistryHost:     "http://localhost:8081",
		Acks:                   "all",
		Timeout:                10,
		UseAuthentication:      false,
		EnableTLS:              true,
		SaslMechanism:          "SCRAM-SHA-512",
		User:                   "kafka",
		Password:               "kafka",
		SchemaRegistryUser:     "",
		SchemaRegistryPassword: "",
		EnableEvents:           true,
		MaxMessageBytes:        0,
		RetryMax:               0,
		DlqTopic:               []string{"account_createorupdate_dlq account_delete_dlq"},
		ConsumerTopic:          []string{"account_createorupdate account_delete"},
		ConsumerGroup:          "account-service",
	}

	kafkaClient := NewKafkaClient(&configKafka)

	assert.NotNil(t, kafkaClient)
}
