package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/mocks"
	"context"
	"errors"
	"testing"

	"github.com/VanessaVallarini/account-toolkit/avros"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAccountService(t *testing.T) {
	mockAccountRepository := mocks.NewIAccountRepository(t)
	accountService := NewAccountService(mockAccountRepository)

	assert.NotNil(t, accountService)
}

func TestCreateOrUpdateAccount(t *testing.T) {
	t.Run("Expect to return success on create or update account", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountCreateOrUpdateEvent

		mockAccountRepository.On("CreateOrUpdate",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
		)

		err := accountService.CreateOrUpdate(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error on create or update account", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountCreateOrUpdateEvent

		mockAccountRepository.On("CreateOrUpdate",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			errors.New("error during create account"),
		)

		err := accountService.CreateOrUpdate(ctx, request)

		assert.Error(t, err)
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("Expect to return success on delete account", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountDeleteEvent
		var reponse models.Account

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			&reponse, nil,
		)

		mockAccountRepository.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
		)

		err := accountService.DeleteAccount(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error on delete account when account does not exists", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountDeleteEvent

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil, nil,
		)

		err := accountService.DeleteAccount(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error on delete account when verify should delete account", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountDeleteEvent

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil, errors.New("error during query get account by email"),
		)

		err := accountService.DeleteAccount(ctx, request)

		assert.Error(t, err)
	})

	t.Run("Expect to return error on delete account", func(t *testing.T) {
		ctx := context.Background()
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockAccountRepository)

		var request avros.AccountDeleteEvent
		var reponse models.Account

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			&reponse, nil,
		)

		mockAccountRepository.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			errors.New("error during query delete account"),
		)

		err := accountService.DeleteAccount(ctx, request)

		assert.Error(t, err)
	})
}
