package main

import (
	"account-consumer-service/cmd/account-consumer-service/health"
	"account-consumer-service/cmd/account-consumer-service/listner"
	"account-consumer-service/cmd/account-consumer-service/server"
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

func main() {

	ctx := context.Background()

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	scylla, err := db.NewScylla(config.Database)
	if err != nil {
		panic(err)
	}
	defer scylla.Close()

	kafkaClient, err := kafka.NewKafkaClient(
		config.Kafka,
	)
	if err != nil {
		panic(err)
	}
	defer kafkaClient.Close()

	kafkaProducer, err := kafkaClient.NewProducer(config.Kafka)
	if err != nil {
		panic(err)
	}

	server := server.NewServer()

	accountRepository := repository.NewAccountRepository(scylla)
	accountService := services.NewAccountService(accountRepository)

	go listner.Start(ctx, config.Kafka, kafkaClient, kafkaProducer, accountService)

	setupHttpServer(server, config)

	utils.Logger.Info("start application")

	health.NewHealthServer()
}

func setupHttpServer(server *server.Server, config *models.Config) {
	go func() {
		server.Start(config)
	}()
}
