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

func TestNewPhoneRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	phoneRepository := NewPhoneRepository(scylla)

	assert.NotNil(t, phoneRepository)
}

func TestCreatePhone(t *testing.T) {
	t.Run("Expect to return success on create phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneDBModel{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := phoneRepository.Create(ctx, p)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneDBModel{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query create phone"),
		)

		err := phoneRepository.Create(ctx, p)

		assert.Error(t, err)
	})
}

func TestUpdatePhone(t *testing.T) {
	t.Run("Expect to return success on update phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneDBModel{}

		scylla.When("Update",
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

		err := phoneRepository.Update(ctx, p)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneDBModel{}

		scylla.When("Update",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query update phone"),
		)

		err := phoneRepository.Update(ctx, p)

		assert.Error(t, err)
	})
}

func TestDeletePhone(t *testing.T) {
	t.Run("Expect to return success on delete phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := phoneRepository.Delete(ctx, p)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query delete phone"),
		)

		err := phoneRepository.Delete(ctx, p)

		assert.Error(t, err)
	})
}

func TestGetByIdPhone(t *testing.T) {
	t.Run("Expect to return success on get phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		address, err := phoneRepository.GetById(ctx, p)

		assert.Nil(t, err)
		assert.NotNil(t, address)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		p := models.PhoneRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get phone by id"),
		)

		address, err := phoneRepository.GetById(ctx, p)

		assert.Error(t, err)
		assert.Nil(t, address)
	})
}

func TestGetAllPhone(t *testing.T) {
	t.Run("Expect to return success on get all phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()
		var pdList []models.PhoneDBModel
		pd := models.PhoneDBModel{
			Id:          randomUUID.String(),
			AreaCode:    "11",
			CountryCode: "55",
			Number:      "964127229",
			UserId:      randomUUID.String(),
		}
		pdList = append(pdList, pd)
		pd = models.PhoneDBModel{
			Id:          randomUUID.String(),
			AreaCode:    "11",
			CountryCode: "55",
			Number:      "964127229",
			UserId:      randomUUID.String(),
		}
		pdList = append(pdList, pd)

		var requestAsMap []map[string]interface{}
		marshalledRequest, _ := json.Marshal(pdList)
		json.Unmarshal(marshalledRequest, &requestAsMap)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			requestAsMap,
			nil,
		)

		phoneList, err := phoneRepository.List(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, phoneList)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		phoneRepository := NewPhoneRepository(scylla)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			nil,
			errors.New("error during query get all phone"),
		)

		phoneList, err := phoneRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, phoneList)
	})
}
