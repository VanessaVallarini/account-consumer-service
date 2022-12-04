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

	shouldCreateAccount := as.shouldCreateAccount(ctx, ae)

	if shouldCreateAccount == true {
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
	}

	return nil
}

func (as *AccountService) shouldCreateAccount(ctx context.Context, ae models.AccountCreateEvent) bool {
	accountGetByEmail := models.AccountRequestByEmail{
		Email: ae.Email,
	}
	account, err := as.getAccountByEmail(ctx, accountGetByEmail)
	if err != nil {
		return false
	}
	if account != nil {
		return false
	}

	accountGetByFullNumber := models.AccountRequestByFullNumber{
		FullNumber: ae.FullNumber,
	}
	account, err = as.getAccountByFullNumber(ctx, accountGetByFullNumber)
	if err != nil {
		return false
	}
	if account != nil {
		return false
	}

	return true
}

func (as *AccountService) getAccountByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	account, err := as.repository.GetByEmail(ctx, a)
	if err != nil {
		utils.Logger.Error("error during get account by email", err)
		return nil, err
	}

	return account, nil
}

func (as *AccountService) getAccountByFullNumber(ctx context.Context, ae models.AccountRequestByFullNumber) (*models.Account, error) {
	account, err := as.repository.GetByFullNumber(ctx, ae)
	if err != nil {
		utils.Logger.Error("error during get account by full number", err)
		return nil, err
	}

	return account, nil
}
