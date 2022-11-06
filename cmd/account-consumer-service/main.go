package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/repository"
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
	ctx := context.Background()
	config := config.NewConfig()
	//Start(ctx, config.Kafka)
	scylla := db.NewScylla(config.Database)

	addressRepository := repository.NewAddressRepository(scylla)

	fmt.Println("PEGANDO TODOS OS ENDEREÇOS...")
	aList, errorx := addressRepository.List(ctx)
	if errorx != nil {
		fmt.Println(errorx)
	}
	for _, a := range aList {
		fmt.Printf("Id:%v. Alias:%s, City:%s, District:%s, PublicPlace:%s, ZipCode:%s \n", a.Id, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode)
	}
	fmt.Println("PEGAMOS TODOS OS ENDEREÇOS!")
	fmt.Println("PEGANDO O ENDEREÇO POR ID...")
	reqAById := models.AddressRequestById{
		Id: aList[0].Id,
	}
	retA, errorx := addressRepository.GetById(ctx, reqAById)
	if errorx != nil {
		fmt.Println(errorx)
	}
	fmt.Printf("Id:%v. Alias:%s, City:%s, District:%s, PublicPlace:%s, ZipCode:%s \n", retA.Id, retA.Alias, retA.City, retA.District, retA.PublicPlace, retA.ZipCode)
	fmt.Println("PEGAMOS O ENDEREÇO POR ID!")

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
