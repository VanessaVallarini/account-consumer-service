package pkg

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/clients"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

const topic = "account"

type IAccountServiceProducer interface {
	Create(ctx context.Context, ae models.AccountCreateRequest) error
}

type AccountServiceProducer struct {
	producer kafka.IProducer
	viaCep   clients.ViaCepApiClient
}

func NewAccountServiceProducer(p kafka.IProducer, v clients.ViaCepApiClient) *AccountServiceProducer {
	return &AccountServiceProducer{
		producer: p,
		viaCep:   v,
	}
}

func (asp *AccountServiceProducer) Create(ctx context.Context, ae models.AccountCreateRequest) error {

	viaCepRequest := models.ViaCepRequest{
		Cep: ae.ZipCode,
	}

	viaCepResponse, err := asp.viaCep.CallViaCepApi(ctx, viaCepRequest)
	if err != nil {
		utils.Logger.Error("error during call via cep api", err)
		return err
	}

	aCreate := models.AccountEvent{
		Alias:       viaCepResponse.Uf,
		City:        viaCepResponse.Localidade,
		District:    viaCepResponse.Bairro,
		Email:       ae.Email,
		FullNumber:  ae.FullNumber,
		Name:        ae.Name,
		PublicPlace: viaCepResponse.Logradouro,
		ZipCode:     ae.ZipCode,
		Command:     models.Create.String(),
	}

	asp.producer.Send(aCreate, topic, models.AccountSubject)
	return nil
}
