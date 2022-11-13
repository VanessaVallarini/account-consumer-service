package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"context"
)

const topic = "account"

type IAccountServiceProducer interface {
	Create(ctx context.Context, ae models.AccountEvent) error
}

type AccountServiceProducer struct {
	producer kafka.IProducer
}

func NewAccountServiceProducer(p kafka.IProducer) *AccountServiceProducer {
	return &AccountServiceProducer{
		producer: p,
	}
}

func (asp *AccountServiceProducer) Create(ctx context.Context, ae models.Account) error {

	aCreate := models.AccountEvent{
		Name:        ae.Name,
		Email:       ae.Email,
		Alias:       ae.Alias,
		City:        ae.City,
		District:    ae.District,
		PublicPlace: ae.PublicPlace,
		ZipCode:     ae.ZipCode,
		CountryCode: ae.CountryCode,
		AreaCode:    ae.AreaCode,
		Number:      ae.Number,
		Command:     models.Create.String(),
	}

	asp.producer.Send(aCreate, topic, models.AccountSubject)
	return nil
}
