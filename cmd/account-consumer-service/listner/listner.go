package listner

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/service/consumer"
	"context"
	"fmt"
)

func Start(ctx context.Context, cfg *models.KafkaConfig) {
	kafkaClient, _ := kafka.NewKafkaClient(cfg)
	err := consumer.NewConsumer(ctx, cfg, kafkaClient)
	if err != nil {
		fmt.Println(err)
	}
}
