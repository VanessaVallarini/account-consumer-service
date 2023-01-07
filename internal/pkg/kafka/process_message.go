package kafka

import (
	"account-consumer-service/internal/pkg/utils"
	"context"

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
	case topic_account_get:
		if err := consumer.getAccount(ctx, message); err != nil {
			return err
		}
	}

	return nil
}

func (consumer *Consumer) createAccount(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account avros.AccountCreateOrUpdateEvent

	if err := consumer.sr.Decode(message.Value, &account, avros.AccountCreateOrUpdateSubject); err != nil {
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
	var account avros.AccountDeleteEvent

	if err := consumer.sr.Decode(message.Value, &account, avros.AccountDeleteSubject); err != nil {
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
	var account avros.AccountGetEvent

	if err := consumer.sr.Decode(message.Value, &account, avros.AccountGetSubject); err != nil {
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

	aSend := avros.AccountGetResponseEvent{
		Email:       account.Email,
		FullNumber:  account.FullNumber,
		Alias:       account.Alias,
		City:        account.City,
		District:    account.District,
		Name:        account.Name,
		PublicPlace: account.PublicPlace,
		Status:      account.Status,
		ZipCode:     account.ZipCode,
	}

	consumer.producer.Send(aSend, topic_account_get_response, avros.AccountResponseSubject)

	if err != nil {
		utils.Logger.Error("error during send msg %v", err)
		return err
	}

	return nil
}
