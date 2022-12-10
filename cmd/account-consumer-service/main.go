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

	"github.com/labstack/echo"
)

func main() {

	ctx := context.Background()

	config := config.NewConfig()

	scylla, err := db.NewScylla(config.Database)
	if err != nil {
		utils.Logger.Error("error during create scylla", err)
	}
	defer scylla.Close()

	accountRepository := repository.NewAccountRepository(scylla)
	accountServiceConsumer := services.NewAccountService(accountRepository)

	kafkaClient, err := kafka.NewKafkaClient(config.Kafka)
	if err != nil {
		utils.Logger.Warn("error during create kafka client")
	}

	go listner.Start(ctx, config.Kafka, accountServiceConsumer, kafkaClient)

	go func() {
		setupHttpServer(config)
	}()

	utils.Logger.Info("start application")

	health.NewHealthServer()
}

func setupHttpServer(config *models.Config) *echo.Echo {
	s := server.NewServer()
	s.Start(config)

	return s.Server
}
