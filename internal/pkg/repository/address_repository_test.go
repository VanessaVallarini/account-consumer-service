package repository

import (
	"account-consumer-service/internal/pkg/db/mocks"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"account-consumer-service/internal/models"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	addressRepository := NewAddressRepository(scylla)

	assert.NotNil(t, addressRepository)
}

func TestCreateAddress(t *testing.T) {
	t.Run("Expect to return success on create address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressDBModel{}

		scylla.When("Insert",
			mock.Any,
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

		err := addressRepository.Create(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressDBModel{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query create address"),
		)

		err := addressRepository.Create(ctx, a)

		assert.Error(t, err)
	})
}

func TestUpdateAddress(t *testing.T) {
	t.Run("Expect to return success on update address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressDBModel{}

		scylla.When("Update",
			mock.Any,
			mock.Any,
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

		err := addressRepository.Update(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressDBModel{}

		scylla.When("Update",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query update address"),
		)

		err := addressRepository.Update(ctx, a)

		assert.Error(t, err)
	})
}

func TestDeleteAddress(t *testing.T) {
	t.Run("Expect to return success on delete address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := addressRepository.Delete(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query delete address"),
		)

		err := addressRepository.Delete(ctx, a)

		assert.Error(t, err)
	})
}

func TestGetByIdAddress(t *testing.T) {
	t.Run("Expect to return success on get address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressRequestById{}

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

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		a := models.AddressRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get address by id"),
		)

		address, err := addressRepository.GetById(ctx, a)

		assert.Error(t, err)
		assert.Nil(t, address)
	})
}

func TestGetAllAddress(t *testing.T) {
	t.Run("Expect to return success on get all address", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()

		var adList []models.AddressDBModel
		ad := models.AddressDBModel{
			Id:          randomUUID.String(),
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
		}
		adList = append(adList, ad)
		ad = models.AddressDBModel{
			Id:          randomUUID.String(),
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
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

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		addressRepository := NewAddressRepository(scylla)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			nil,
			errors.New("error during query get all address"),
		)

		addressList, err := addressRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, addressList)
	})
}
