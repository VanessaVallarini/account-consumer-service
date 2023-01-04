package kafka

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"
	"context"

	"github.com/Shopify/sarama"
)

const (
	topic_create = "account_create"
	topic_update = "account_update"
	topic_delete = "account_delete"
)

func (consumer *Consumer) processMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	switch message.Topic {
	case topic_create:
		if err := consumer.createAccount(ctx, message); err != nil {
			return err
		}
	case topic_update:
		if err := consumer.updateAccount(ctx, message); err != nil {
			return err
		}
	case topic_delete:
		if err := consumer.deleteAccount(ctx, message); err != nil {
			return err
		}
	}

	return nil
}

func (consumer *Consumer) createAccount(ctx context.Context, message *sarama.ConsumerMessage) error {

	return nil
}

func (consumer *Consumer) updateAccount(ctx context.Context, message *sarama.ConsumerMessage) error {

	return nil
}

func (consumer *Consumer) deleteAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var ac models.AccountDeleteEvent

	if err := consumer.sr.Decode(message.Value, &ac, models.AccountDeleteSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	//if err := consumer.accountServiceConsumer.DeleteAccount(ctx, ac); err != nil {
	//return err
	//}

	return nil
}
