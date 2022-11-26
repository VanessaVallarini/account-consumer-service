package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/services"
	"context"
	"fmt"
)

func Start(ctx context.Context, cfg *models.KafkaConfig, asc *services.AccountService) {
	kafkaClient, _ := kafka.NewKafkaClient(cfg)
	err := kafka.NewConsumer(ctx, cfg, kafkaClient, asc)
	if err != nil {
		fmt.Println(err)
	}
}
