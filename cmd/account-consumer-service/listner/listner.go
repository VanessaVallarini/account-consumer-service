package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

func Start(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *kafka.KafkaClient, kafkaProducer kafka.IKafkaProducer, accountService *services.AccountService) {
	err := kafka.NewConsumer(ctx, cfg, kafkaClient, kafkaProducer, accountService)
	if err != nil {
		utils.Logger.Errorf("Error consumer msg: %v", err)
	}
}
