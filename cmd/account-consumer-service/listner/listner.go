package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"context"
	"fmt"
)

func Start(ctx context.Context, cfg *models.KafkaConfig) {
	kafkaClient, _ := kafka.NewKafkaClient(cfg)
	err := kafka.NewConsumer(ctx, cfg, kafkaClient)
	if err != nil {
		fmt.Println(err)
	}
}
