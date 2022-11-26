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
	repository repository.IAccountRepository
}

func NewAccountService(repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		repository: repo,
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

	accountCreate := models.Account{
		Alias:       ae.Alias,
		City:        ae.City,
		District:    ae.District,
		Email:       ae.Email,
		FullNumber:  ae.FullNumber,
		Name:        ae.Name,
		PublicPlace: ae.PublicPlace,
		ZipCode:     ae.ZipCode,
	}

	err = as.repository.Create(ctx, accountCreate)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}

	return nil
}

func (as *AccountService) shouldCreateAccount(ctx context.Context, ae models.AccountEvent) (bool, error) {
	accountRequestBy := models.AccountRequestBy{
		Id:         "",
		Email:      ae.Email,
		FullNumber: ae.FullNumber,
	}

	account, err := as.getAccountBy(ctx, accountRequestBy)
	if err != nil {
		return false, err
	}
	if account != nil {
		utils.Logger.Error("account already exists", err)
		return false, err
	}

	return true, nil
}

func (as *AccountService) getAccountBy(ctx context.Context, a models.AccountRequestBy) (*models.Account, error) {

	account, err := as.repository.GetBy(ctx, a)
	if err != nil {
		utils.Logger.Error("error during get account by", err)
		return nil, err
	}

	return account, nil
}
