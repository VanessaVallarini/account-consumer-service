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

func TestNewAccountRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	userRepository := NewAccountRepository(scylla)

	assert.NotNil(t, userRepository)
}

func TestCreateAccount(t *testing.T) {
	t.Run("Expect to return success on create account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.Account{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
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

		err := accountRepository.Create(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query on create account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.Account{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
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
			errors.New("error during query create account"),
		)

		err := accountRepository.Create(ctx, a)

		assert.Error(t, err)
	})
}

func TestGetByIdAccount(t *testing.T) {
	t.Run("Expect to return success on get account by id", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		account, err := accountRepository.GetById(ctx, a)

		assert.Nil(t, err)
		assert.NotNil(t, account)
	})

	t.Run("Expect to return error during query on get account by id", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by id"),
		)

		account, err := accountRepository.GetById(ctx, a)

		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestGetByEmailAccount(t *testing.T) {
	t.Run("Expect to return success on get account by email", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestByEmail{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		account, err := accountRepository.GetByEmail(ctx, a)

		assert.Nil(t, err)
		assert.NotNil(t, account)
	})

	t.Run("Expect to return error during query on get account by email", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestByEmail{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by email"),
		)

		account, err := accountRepository.GetByEmail(ctx, a)

		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestGetByPhoneAccount(t *testing.T) {
	t.Run("Expect to return success on get account by phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestByPhone{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		account, err := accountRepository.GetByPhone(ctx, a)

		assert.Nil(t, err)
		assert.NotNil(t, account)
	})

	t.Run("Expect to return error during query on get account by phone", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestByPhone{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by phone"),
		)

		account, err := accountRepository.GetByPhone(ctx, a)

		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestGetAllAccount(t *testing.T) {
	t.Run("Expect to return success on get all account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()
		var aList []models.Account
		a := models.Account{
			Id:          randomUUID.String(),
			Name:        "Lorem",
			Email:       "lorem@email.com",
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
			CountryCode: "55",
			AreaCode:    "11",
			Number:      "964127229",
		}
		aList = append(aList, a)
		aList = append(aList, a)

		var requestAsMap []map[string]interface{}
		marshalledRequest, _ := json.Marshal(aList)
		json.Unmarshal(marshalledRequest, &requestAsMap)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			requestAsMap,
			nil,
		)

		accountList, err := accountRepository.List(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, accountList)
	})

	t.Run("Expect to return error during query on get all account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			nil,
			errors.New("error during query get all account"),
		)

		accountList, err := accountRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, accountList)
	})
}

func TestUpdateAccount(t *testing.T) {
	t.Run("Expect to return success on update account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.Account{}

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
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := accountRepository.Update(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query on update account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.Account{}

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
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query update account"),
		)

		err := accountRepository.Update(ctx, a)

		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Expect to return success on delete account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := accountRepository.Delete(ctx, a)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query on delete account", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		p := models.AccountRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query delete account"),
		)

		err := accountRepository.Delete(ctx, p)

		assert.Error(t, err)
	})
}
