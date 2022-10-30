package consumer

import (
	"account-consumer-service/internal/entities"
	"context"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// data holds in memory the messages until the bufferMaxSize
var data []entities.AccountEvent

var bufferMaxSize = viper.GetInt("BUFFER_MAX_SIZE")

func (consumer *Consumer) processMessage(ctx context.Context, message *sarama.ConsumerMessage, session sarama.ConsumerGroupSession) error {
	zap.S().Infof("Message claimed: topic = %s, partition = %v, offset = %v", message.Topic, message.Partition, message.Offset)

	var msg entities.AccountEvent
	err := consumer.sr.Decode(message.Value, &msg, entities.AccountSubject)
	if err != nil {
		zap.S().Infof("could not decode value %v", err)
		return err
	}

	data = append(data, entities.AccountEvent{
		Id: msg.Id,

		UserId: msg.UserId,
		Name:   msg.Name,
		Email:  msg.Email,

		AddressId:   msg.AddressId,
		Alias:       msg.Alias,
		City:        msg.City,
		District:    msg.District,
		PublicPlace: msg.PublicPlace,
		ZipCode:     msg.ZipCode,

		PhoneId:     msg.PhoneId,
		CountryCode: msg.CountryCode,
		AreaCode:    msg.AreaCode,
		Number:      msg.Number,
	})

	if len(data) < bufferMaxSize {
		return nil
	}

	data = make([]entities.AccountEvent, 0)
	session.Commit()
	return nil
}
