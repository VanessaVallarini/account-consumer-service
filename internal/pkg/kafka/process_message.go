package kafka

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"
	"context"

	"github.com/Shopify/sarama"
)

func (consumer *Consumer) processMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	var ac models.AccountCreateEvent

	if err := consumer.sr.Decode(message.Value, &ac, models.AccountCreateSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	if err := consumer.accountServiceConsumer.CreateAccount(ctx, ac); err != nil {
		utils.Logger.Error("error during create account")
		return err
	}

	return nil
}
