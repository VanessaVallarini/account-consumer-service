package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db/mocks"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	addressRepository := NewAddressRepository(scylla)

	assert.NotNil(t, addressRepository)
}

func TestCreate(t *testing.T) {
	t.Run("Expect to return success on create address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.Address{
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
		}

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
	})

	t.Run("Expect to return error during insert query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.Address{
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
		}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during insert query"),
		)

		err := addressRepository.Insert(ctx, a)

		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Expect to return success on get address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()

		a := models.AddressRequestById{
			Id: randomUUID.String(),
		}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		address, err := addressRepository.GetById(ctx, a)

		assert.Nil(t, err)
		assert.NotNil(t, address)
	})

	t.Run("Expect to return error on get address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()

		a := models.AddressRequestById{
			Id: randomUUID.String(),
		}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during insert query"),
		)

		address, err := addressRepository.GetById(ctx, a)

		assert.Error(t, err)
		assert.Nil(t, address)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Expect to return success on get all address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()

		var adList []models.AddressDbo
		ad := models.AddressDbo{
			Id:           randomUUID,
			Alias:        "SP",
			City:         "São Paulo",
			District:     "Sé",
			Public_place: "Praça da Sé",
			Zip_code:     "01001-000",
		}
		adList = append(adList, ad)
		ad = models.AddressDbo{
			Id:           randomUUID,
			Alias:        "SP",
			City:         "São Paulo",
			District:     "Sé",
			Public_place: "Praça da Sé",
			Zip_code:     "01001-000",
		}
		adList = append(adList, ad)

		var requestAsMap []map[string]interface{}
		marshalledRequest, _ := json.Marshal(adList)
		json.Unmarshal(marshalledRequest, &requestAsMap)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			requestAsMap,
			nil,
		)

		addressList, err := addressRepository.List(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, addressList)
	})

	t.Run("Expect to return error on get all address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			nil,
			errors.New("error during insert query"),
		)

		addressList, err := addressRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, addressList)
	})
}
