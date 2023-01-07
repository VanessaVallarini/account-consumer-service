package kafka

import (
	"account-consumer-service/internal/pkg/utils"
	"context"
	"errors"

	"github.com/Shopify/sarama"
	"github.com/VanessaVallarini/account-toolkit/avros"
)

func (consumer *Consumer) processMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	switch message.Topic {
	case topic_account_createorupdate:
		if err := consumer.createAccount(ctx, message); err != nil {
			return err
		}
	case topic_account_delete:
		if err := consumer.deleteAccount(ctx, message); err != nil {
			return err
		}
	default:
		return errors.New("Invalid topic.")
	}

	return nil
}

func (consumer *Consumer) createAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account avros.AccountCreateOrUpdateEvent

	if err := consumer.sr.Decode(message.Value, &account, avros.AccountCreateOrUpdateSubject); err != nil {
		utils.Logger.Error("Error during decode message consumer kafka on create account. Details: %v", err)
		return err
	}

	if err := consumer.accountService.CreateOrUpdate(ctx, account); err != nil {
		return err
	}

	return nil
}

func (consumer *Consumer) deleteAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account avros.AccountDeleteEvent

	if err := consumer.sr.Decode(message.Value, &account, avros.AccountDeleteSubject); err != nil {
		utils.Logger.Error("Error during decode message consumer kafka on delete account. Details: %v", err)
		return err
	}

	if err := consumer.accountService.DeleteAccount(ctx, account); err != nil {
		return err
	}

	return nil
}
