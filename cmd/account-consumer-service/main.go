package main

import (
	api "account-consumer-service/api/account"
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
	"account-consumer-service/pkg"
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

	kafkaProducer, err := kafkaClient.NewProducer()
	if err != nil {
		utils.Logger.Warn("error during kafka producer")
	}

	accountServiceProducer := pkg.NewAccountServiceProducer(*kafkaProducer)

	go func() {
		setupHttpServer(accountServiceProducer, config)
	}()

	go listner.Start(ctx, config.Kafka, accountServiceConsumer)

	utils.Logger.Info("start application")

	health.NewHealthServer()

}

func setupHttpServer(asp *pkg.AccountServiceProducer, config *models.Config) *echo.Echo {

	accountApi := api.NewAccountApi(asp)
	s := server.NewServer()
	accountApi.Register(s.Server)

	s.Start(config)

	return s.Server
}
