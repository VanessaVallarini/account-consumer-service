package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/service/consumer"
	"context"
	"fmt"
)

func Start(ctx context.Context, cfg *models.KafkaConfig) {
	kafkaClient, _ := kafka.NewKafkaClient(cfg)
	err := consumer.NewConsumer(ctx, cfg, kafkaClient)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//var err error

	ctx := context.Background()
	config := config.NewConfig()
	Start(ctx, config.Kafka)
	//scylla := scylla.NewScylla(config.Database)

	/* kafkaClient, _ := kafka.NewKafkaClient(config.Kafka)
	p, _ := kafkaClient.NewProducer()
	aCreate := models.AccountEvent{
		Name:        "name",
		Email:       "email",
		Alias:       "alias",
		City:        "city",
		District:    "district",
		PublicPlace: "public_place",
		ZipCode:     "zip_code",
		CountryCode: "country_code",
		AreaCode:    "area_code",
		Number:      "number",
	}
	p.Send(aCreate, config.Kafka.ConsumerTopic, models.AccountSubject) */
}
