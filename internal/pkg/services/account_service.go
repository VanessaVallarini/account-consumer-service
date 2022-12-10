package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/utils"
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

	accountCreate := models.AccountCreate{
		Alias:       ae.Alias,
		City:        ae.City,
		District:    ae.District,
		Email:       ae.Email,
		FullNumber:  ae.FullNumber,
		Name:        ae.Name,
		PublicPlace: ae.PublicPlace,
		ZipCode:     ae.ZipCode,
	}

	err := as.repository.Create(ctx, accountCreate)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}

	return nil
}
