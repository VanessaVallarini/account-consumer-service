package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/repository"
	"context"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, ae models.AccountCreateEvent) error
}

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		repository: repo,
	}
}

func (as *AccountService) CreateAccount(ctx context.Context, ae models.AccountCreateEvent) error {

	return nil
}

func (as *AccountService) UpdateAccount(ctx context.Context, ae models.AccountUpdateEvent) error {

	//err := as.repository.Update(ctx, accountUpdate)
	//if err != nil {
	//utils.Logger.Error("error during create account", err)
	//return err
	//}

	return nil
}
