package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, ae models.AccountEvent) error
}

type AccountService struct {
	registry repository.Registry
}

func NewAccountService(r repository.Registry) *AccountService {
	return &AccountService{
		registry: r,
	}
}

func (as *AccountService) CreateAccount(ctx context.Context, ae models.AccountEvent) error {
	ok, err := as.shouldCreateAccount(ctx, ae)
	if err != nil {
		return err
	}
	if !ok {
		return err
	}

	accountRepository := as.registry.AccountRepository()

	accountCreate := models.Account{
		Name:        ae.Name,
		Email:       ae.Email,
		Alias:       ae.Alias,
		City:        ae.City,
		District:    ae.District,
		PublicPlace: ae.PublicPlace,
		ZipCode:     ae.ZipCode,
		CountryCode: ae.CountryCode,
		AreaCode:    ae.AreaCode,
		Number:      ae.Number,
	}

	err = accountRepository.Create(ctx, accountCreate)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}

	return nil
}

func (as *AccountService) shouldCreateAccount(ctx context.Context, ae models.AccountEvent) (bool, error) {
	accountRequestByEmail := models.AccountRequestByEmail{
		Email: ae.Email,
	}

	account, err := as.getAccountByEmail(ctx, accountRequestByEmail)
	if err != nil {
		return false, err
	}
	if account != nil {
		utils.Logger.Error("account already exists", err)
		return false, err
	}

	accountRequestByPhone := models.AccountRequestByPhone{
		CountryCode: ae.CountryCode,
		AreaCode:    ae.AreaCode,
		Number:      ae.Number,
	}

	account, err = as.getAccountByPhone(ctx, accountRequestByPhone)
	if err != nil {
		return false, err
	}
	if account != nil {
		utils.Logger.Error("account already exists", err)
		return false, err
	}

	return true, nil
}

func (as *AccountService) getAccountByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	accountRepository := as.registry.AccountRepository()

	account, err := accountRepository.GetByEmail(ctx, a)
	if err != nil {
		utils.Logger.Error("error during get account by email", err)
		return nil, err
	}

	return account, nil
}

func (as *AccountService) getAccountByPhone(ctx context.Context, a models.AccountRequestByPhone) (*models.Account, error) {
	accountRepository := as.registry.AccountRepository()

	account, err := accountRepository.GetByPhone(ctx, a)
	if err != nil {
		utils.Logger.Error("error during get account by phone", err)
		return nil, err
	}

	return account, nil
}
