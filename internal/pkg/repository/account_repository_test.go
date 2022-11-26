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
	accountRepository := NewAccountRepository(scylla)

	assert.NotNil(t, accountRepository)
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
		).Return(
			errors.New("error during query create account"),
		)

		err := accountRepository.Create(ctx, a)

		assert.Error(t, err)
	})
}

func TestGetByAccount(t *testing.T) {
	t.Run("Expect to return success on get account by", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestBy{}

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

		account, err := accountRepository.GetBy(ctx, a)

		assert.Nil(t, err)
		assert.NotNil(t, account)
	})

	t.Run("Expect to return success during query on get account by and account not exist", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestBy{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("not found"),
		)

		account, err := accountRepository.GetBy(ctx, a)

		assert.Nil(t, err)
		assert.Nil(t, account)
	})

	t.Run("Expect to return error during query on get account by", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		accountRepository := NewAccountRepository(scylla)

		a := models.AccountRequestBy{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by"),
		)

		account, err := accountRepository.GetBy(ctx, a)

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
			Alias:       "SP",
			City:        "São Paulo",
			District:    "Sé",
			Email:       "lorem@email.com",
			FullNumber:  "5511964127229",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			ZipCode:     "01001-000",
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
