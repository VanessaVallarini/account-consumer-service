package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/entities"
	"account-consumer-service/internal/pkg/kafka"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/riferrei/srclient"
)

const (
	UserSubject = "com.account.producer"
	UserAvro    = `{
		"type": "record",
		"name": "UserAccount",
		"namespace": "com.account.producer",
		"fields": [
			{ "name": "name", "type": "string" },
			{ "name": "email", "type": "string" }
		   ]
	   }`
)

func main() {

	//ctx := context.Background()
	config := config.NewConfig()
	//scylla := scylla.NewScylla(config.Database)

	kafkaClient, _ := kafka.NewKafkaClient(config.Kafka)
	ret, err := kafkaClient.SchemaRegistry.CreateSchema(UserSubject, UserAvro, srclient.Avro)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
	Producer(config, kafkaClient)

}

func Producer(config *entities.Config, client *kafka.KafkaClient) {
	producer, err := sarama.NewSyncProducerFromClient(client.Client)
	if err != nil {
		fmt.Println(err)
	}

	uCreate := entities.User{
		Name:  "teste_name",
		Email: "teste_email",
	}

	msgEncoder, err := client.SchemaRegistry.Encode(uCreate, entities.UserSubject)
	if err != nil {
		fmt.Println(err)
	}

	msg := sarama.ProducerMessage{
		Topic:     config.Kafka.ConsumerTopic,
		Key:       sarama.ByteEncoder(time.Now().String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}
	producer.SendMessage(&msg)

}
