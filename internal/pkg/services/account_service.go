package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"fmt"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, ae models.AccountCreateEvent) error
	GetAccountByEmail(ctx context.Context, ae models.AccountRequestByEmail) (*models.Account, error)
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

	accountGetByEmail := models.AccountRequestByEmail{
		Email: ae.Email,
	}
	abe, err := as.getAccountByEmail(ctx, accountGetByEmail)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}
	fmt.Println(abe)

	accountGetByFullNumber := models.AccountRequestByFullNumber{
		FullNumber: ae.FullNumber,
	}
	abfn, err := as.getAccountByFullNumber(ctx, accountGetByFullNumber)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}
	fmt.Println(abfn)

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

	err = as.repository.Create(ctx, accountCreate)
	if err != nil {
		utils.Logger.Error("error during create account", err)
		return err
	}

	return nil
}

func (as *AccountService) getAccountByEmail(ctx context.Context, ae models.AccountRequestByEmail) (*models.Account, error) {
	account, err := as.repository.GetByEmail(ctx, ae)
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
