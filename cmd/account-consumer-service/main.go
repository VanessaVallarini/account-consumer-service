package main

import (
	api "account-consumer-service/api/account"
	"account-consumer-service/cmd/account-consumer-service/listner"
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"account-consumer-service/pkg"
	"context"
)

func main() {
	ctx := context.Background()

	config := config.NewConfig()

	scylla, err := db.NewScylla(config.Database)
	if err != nil {
		utils.Logger.Warn("error during create scylla", err)
	}
	defer scylla.Close()

	accountRepository := repository.NewAccountRepository(scylla)
	accountServiceConsumer := services.NewAccountService(accountRepository)

	go func() {

		kafkaClient, err := kafka.NewKafkaClient(config.Kafka)
		if err != nil {
			utils.Logger.Warn("error during kafka client")
		}

		kafkaProducer, err := kafkaClient.NewProducer()
		if err != nil {
			utils.Logger.Warn("error during kafka producer")
		}

		accountServiceProducer := pkg.NewAccountServiceProducer(*kafkaProducer)
		accountApi := api.NewAccountApi(accountServiceProducer)
		server := utils.NewServer()
		accountApi.Register(server.Server)

		listner.Start(ctx, config.Kafka, accountServiceConsumer)

		server.Start(config)
	}()

	utils.Logger.Info("start application")

	utils.NewHealthServer()

}
