package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/services"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

func Start(ctx context.Context, cfg *models.KafkaConfig, asc *services.AccountService) {
	kafkaClient, _ := kafka.NewKafkaClient(cfg)
	err := kafka.NewConsumer(ctx, cfg, kafkaClient, asc)
	if err != nil {
		utils.Logger.Error("Error consumer msg: %v", err)
	}
}
