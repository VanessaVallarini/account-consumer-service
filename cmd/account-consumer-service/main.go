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

	scylla := db.NewScylla(config.Database)
	defer scylla.Close()

	kafkaClient := kafka.NewKafkaClient(config.Kafka)

	accountRepository := repository.NewAccountRepository(scylla)
	accountService := services.NewAccountService(accountRepository)

	go listner.Start(ctx, config.Kafka, kafkaClient, accountService)

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
