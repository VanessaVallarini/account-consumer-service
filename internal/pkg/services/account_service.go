package services

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"time"

	"github.com/VanessaVallarini/account-toolkit/avros"
)

type IAccountService interface {
	CreateOrUpdateAccount(ctx context.Context, ae avros.AccountCreateOrUpdateEvent) error
	DeleteAccount(ctx context.Context, ade avros.AccountDeleteEvent)
}

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		repository: repo,
	}
}

func (service *AccountService) CreateOrUpdate(ctx context.Context, ace avros.AccountCreateOrUpdateEvent) error {

	account := models.Account{
		Email:       ace.Email,
		FullNumber:  ace.FullNumber,
		Alias:       ace.Alias,
		City:        ace.City,
		DateTime:    time.Now().String(),
		District:    ace.District,
		Name:        ace.Name,
		PublicPlace: ace.PublicPlace,
		Status:      models.AccountStatusString(ace.Status).String(),
		ZipCode:     ace.ZipCode,
	}

	err := service.repository.CreateOrUpdate(ctx, account)
	if err != nil {
		utils.Logger.Errorf("error during create account", err)
		return err
	}

	return nil
}

func (service *AccountService) DeleteAccount(ctx context.Context, ade avros.AccountDeleteEvent) error {
	request := models.AccountRequestByEmail{
		Email: ade.Email,
	}

	shouldCreateAccount, err := service.shouldDeleteAccount(ctx, request)
	if err != nil {
		utils.Logger.Errorf("error during verify should update account", err)
		return err
	}

	if shouldCreateAccount {
		err := service.repository.Delete(ctx, request)
		if err != nil {
			utils.Logger.Errorf("error during delete account", err)
			return err
		}
	} else {
		return err
	}

	return nil
}

func (service *AccountService) shouldDeleteAccount(ctx context.Context, request models.AccountRequestByEmail) (bool, error) {
	accountRespByEmailAndFullNumber, err := service.repository.GetByEmail(ctx, request)
	if accountRespByEmailAndFullNumber == nil {
		utils.Logger.Errorf("account does not exist", err)
		return false, nil
	}

	return true, nil
}
