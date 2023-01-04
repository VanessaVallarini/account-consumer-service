package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

func Start(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *kafka.KafkaClient, accountService *services.AccountService) {
	err := kafka.NewConsumer(ctx, cfg, kafkaClient, accountService)
	if err != nil {
		utils.Logger.Error("Error consumer msg: %v", err)
	}
}
