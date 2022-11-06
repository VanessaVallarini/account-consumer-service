package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db/mocks"
	"context"
	"errors"
	"testing"

	"github.com/maraino/go-mock"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	addressRepository := NewAddressRepository(scylla)

	assert.NotNil(t, addressRepository)
}

func TestCreateAddressReturnSuccess(t *testing.T) {
	ctx := context.Background()
	scylla := mocks.NewScylla()
	addressRepository := NewAddressRepository(scylla)

	a := createAddressParams()

	scylla.When("Insert",
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
	).Return(
		nil,
	)

	err := addressRepository.Insert(ctx, a)

	assert.Nil(t, err)
}

func TestCreateAddressReturnError(t *testing.T) {
	ctx := context.Background()
	scylla := mocks.NewScylla()
	addressRepository := NewAddressRepository(scylla)

	a := createAddressParams()

	scylla.When("Insert",
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
		mock.Any,
	).Return(
		errors.New("error on querying"),
	)

	err := addressRepository.Insert(ctx, a)

	assert.Error(t, err)
}

func createAddressParams() models.Address {
	return models.Address{
		Alias:       "SP",
		City:        "São Paulo",
		District:    "Sé",
		PublicPlace: "Praça da Sé",
		ZipCode:     "01001-000",
	}
}
