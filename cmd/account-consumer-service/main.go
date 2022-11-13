package main

import (
	"account-consumer-service/internal/config"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
)

func main() {
	//ctx := context.Background()

	config := config.NewConfig()

	scylla, err := db.NewScylla(config.Database)
	if err != nil {
		utils.Logger.Warn("error during create scylla", err)
	}
	defer scylla.Close()

	utils.Logger.Info("start application")

	utils.NewHealthServer()

	server := utils.NewServer()
	if server == nil {
		utils.Logger.Warn("error during create scylla", err)
	}
	//kafkaClient, err := kafka.NewKafkaClient(config.Kafka)

	//kafkaProducer, err := kafkaClient.NewProducer()
	server.Start(config)
}
