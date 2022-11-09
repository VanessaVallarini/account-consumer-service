package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
)

func main() {
	//ctx := context.Background()
	config := config.NewConfig()

	//scylla := db.NewScylla(config.Database)

	kafkaClient, _ := kafka.NewKafkaClient(config.Kafka)
	p, _ := kafkaClient.NewProducer()
	aCreate := models.AccountEvent{
		Id:          "id",
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
		Command:     "insert",
	}
	p.Send(aCreate, config.Kafka.ConsumerTopic, models.AccountSubject)

	//listner.Start(ctx, config.Kafka)
}
