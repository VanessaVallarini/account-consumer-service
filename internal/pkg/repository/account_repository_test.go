package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountRepository(t *testing.T) {
	mockScylla := mocks.NewIScylla(t)
	accountRepository := NewAccountRepository(mockScylla)

	assert.NotNil(t, accountRepository)
}

func TestCreateOrUpdateAccount(t *testing.T) {
	t.Run("Expect to return success on create or update account", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		var account models.Account

		mockScylla.On("Insert",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		err := accountRepository.CreateOrUpdate(ctx, account)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query on create or update account", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.Account{}

		mockScylla.On("Insert",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during query create account"),
		)

		err := accountRepository.CreateOrUpdate(ctx, request)

		assert.Error(t, err)
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("Expect to return success on delete account", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.AccountRequestByEmail{}

		mockScylla.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		err := accountRepository.Delete(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query on delete account", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.AccountRequestByEmail{}

		mockScylla.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during query delete account"),
		)

		err := accountRepository.Delete(ctx, request)

		assert.Error(t, err)
	})
}

func TestGetAccountByEmail(t *testing.T) {
	t.Run("Expect to return success on get account by email", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.AccountRequestByEmail{}

		mockScylla.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Expect to return error during query on get account by email and account not exist", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.AccountRequestByEmail{}

		mockScylla.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("not found"),
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.Nil(t, response)
	})

	t.Run("Expect to return error during query on get account by email", func(t *testing.T) {
		ctx := context.Background()
		mockScylla := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(mockScylla)

		request := models.AccountRequestByEmail{}

		mockScylla.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during query get account by email"),
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
