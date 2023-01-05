package kafka

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"time"

	"github.com/Shopify/sarama"
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
	case topic_account_get:
		if err := consumer.getAccount(ctx, message); err != nil {
			return err
		}
	}

	return nil
}

func (consumer *Consumer) createAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account models.AccountCreateOrUpdateEvent

	if err := consumer.sr.Decode(message.Value, &account, models.AccountCreateOrUpdateSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	if err := consumer.accountService.CreateOrUpdate(ctx, account); err != nil {
		return err
	}

	if err := consumer.sendAccount(ctx, account.Email); err != nil {
		return err
	}

	return nil
}

func (consumer *Consumer) deleteAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account models.AccountDeleteEvent

	if err := consumer.sr.Decode(message.Value, &account, models.AccountDeleteSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	if err := consumer.accountService.DeleteAccount(ctx, account); err != nil {
		return err
	}

	if err := consumer.sendAccount(ctx, account.Email); err != nil {
		return err
	}

	return nil
}

func (consumer *Consumer) getAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account models.AccountGetEvent

	if err := consumer.sr.Decode(message.Value, &account, models.AccountGetSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	if err := consumer.sendAccount(ctx, account.Email); err != nil {
		return err
	}

	return nil
}

func (consumer *Consumer) sendAccount(ctx context.Context, email string) error {

	account, err := consumer.accountService.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	msgEncoder, err := consumer.sr.Encode(account, models.AccountGetResponseSubject)
	if err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic_account_get_response,
		Key:       sarama.ByteEncoder(time.Now().String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}

	partition, offset, err := consumer.producer.SendMessage(msg)
	if err != nil {
		utils.Logger.Error("error during send msg. partition: %v, offset: %v", msg, partition, offset)
		return err
	}

	return nil
}
