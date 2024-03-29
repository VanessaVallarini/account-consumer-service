package kafka

import (
	"account-consumer-service/internal/models"
	"testing"

	"github.com/VanessaVallarini/account-toolkit/avros"
	"github.com/stretchr/testify/assert"
)

func TestKafkaProducer(t *testing.T) {
	configKafka := &models.KafkaConfig{
		ClientId:               "account-producer-service",
		Hosts:                  []string{"localhost:9092"},
		UseAuthentication:      false,
		SaslMechanism:          "SCRAM-SHA-512",
		EnableTLS:              true,
		SchemaRegistryHost:     "http://localhost:8086",
		User:                   "",
		Password:               "",
		SchemaRegistryUser:     "",
		SchemaRegistryPassword: "",
	}

	kafkaClient, _ := NewKafkaClient(configKafka)

	producer, _ := kafkaClient.NewProducer(configKafka)

	assert.NotNil(t, producer)
}

func TestKafkaProducerSendMsgDlqError(t *testing.T) {
	t.Run("Expect to return error during send msg in DLQ to create account", func(t *testing.T) {
		configKafka := &models.KafkaConfig{
			ClientId:               "account-producer-service",
			Hosts:                  []string{"localhost:9092"},
			UseAuthentication:      false,
			SaslMechanism:          "SCRAM-SHA-512",
			EnableTLS:              true,
			SchemaRegistryHost:     "http://localhost:8086",
			User:                   "",
			Password:               "",
			SchemaRegistryUser:     "",
			SchemaRegistryPassword: "",
		}

		kafkaClient, _ := NewKafkaClient(configKafka)

		producer, _ := kafkaClient.NewProducer(configKafka)

		account := models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}

		err := producer.Send(account, topic_account_createorupdate, avros.AccountCreateOrUpdateSubject)

		assert.Error(t, err)
	})

	t.Run("Expect to return error during send msg in DLQ to delete account", func(t *testing.T) {
		configKafka := &models.KafkaConfig{
			ClientId:               "account-producer-service",
			Hosts:                  []string{"localhost:9092"},
			UseAuthentication:      false,
			SaslMechanism:          "SCRAM-SHA-512",
			EnableTLS:              true,
			SchemaRegistryHost:     "http://localhost:8086",
			User:                   "",
			Password:               "",
			SchemaRegistryUser:     "",
			SchemaRegistryPassword: "",
		}

		kafkaClient, _ := NewKafkaClient(configKafka)

		producer, _ := kafkaClient.NewProducer(configKafka)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		err := producer.Send(request, topic_account_delete, avros.AccountDeleteSubject)

		assert.Error(t, err)
	})
}

func TestKafkaProducerSendMsgDlqSuccess(t *testing.T) {
	t.Run("Expect to return success during send msg in DLQ to create account", func(t *testing.T) {
		configKafka := &models.KafkaConfig{
			ClientId:               "account-producer-service",
			Hosts:                  []string{"localhost:9092"},
			UseAuthentication:      false,
			SaslMechanism:          "SCRAM-SHA-512",
			EnableTLS:              true,
			SchemaRegistryHost:     "http://localhost:8086",
			User:                   "",
			Password:               "",
			SchemaRegistryUser:     "",
			SchemaRegistryPassword: "",
		}

		kafkaClient, _ := NewKafkaClient(configKafka)

		producer, _ := kafkaClient.NewProducer(configKafka)

		account := avros.AccountCreateOrUpdateEvent{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}

		err := producer.Send(account, topic_account_createorupdate, avros.AccountCreateOrUpdateSubject)

		assert.Nil(t, err)
	})

	t.Run("Expect to return success during send msg in DLQ to delete account", func(t *testing.T) {
		configKafka := &models.KafkaConfig{
			ClientId:               "account-producer-service",
			Hosts:                  []string{"localhost:9092"},
			UseAuthentication:      false,
			SaslMechanism:          "SCRAM-SHA-512",
			EnableTLS:              true,
			SchemaRegistryHost:     "http://localhost:8086",
			User:                   "",
			Password:               "",
			SchemaRegistryUser:     "",
			SchemaRegistryPassword: "",
		}

		kafkaClient, _ := NewKafkaClient(configKafka)

		producer, _ := kafkaClient.NewProducer(configKafka)

		request := avros.AccountDeleteEvent{
			Email: "lorem1@email.com",
		}

		err := producer.Send(request, topic_account_delete, avros.AccountDeleteSubject)

		assert.Nil(t, err)
	})
}
