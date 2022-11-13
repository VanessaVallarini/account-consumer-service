package main

import (
	api "account-consumer-service/api/account"
	"account-consumer-service/cmd/account-consumer-service/listner"
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/utils"
	"account-consumer-service/internal/services"
	"context"

	"github.com/labstack/echo"
)

func main() {
	ctx := context.Background()

	config := config.NewConfig()

	go func() {
		setupHttpServer(config)
	}()

	utils.Logger.Info("start application")

	utils.NewHealthServer()

	listner.Start(ctx, config.Kafka)
}

func setupHttpServer(cfg *models.Config) *echo.Echo {

	scylla, err := db.NewScylla(cfg.Database)
	if err != nil {
		utils.Logger.Warn("error during create scylla", err)
	}
	defer scylla.Close()

	server := utils.NewServer()
	if server == nil {
		utils.Logger.Warn("error during create scylla")
	}

	kafkaClient, err := kafka.NewKafkaClient(cfg.Kafka)
	if server == nil {
		utils.Logger.Warn("error during create scylla")
	}

	kafkaProducer, err := kafkaClient.NewProducer()

	//registry := repository.NewRegistry(scylla)
	//repository := repository.NewAccountRepository(scylla)

	//accountService := services.NewAccountService(*registry)
	accountServiceProducer := services.NewAccountServiceProducer(*kafkaProducer)

	accountApi := api.NewAccountApi(*accountServiceProducer)
	accountApi.Register(server.Server)

	server.Start(cfg)

	return server.Server
}
