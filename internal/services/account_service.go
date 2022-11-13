package service

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/kafka"
	"account-consumer-service/internal/pkg/repository"
	"account-consumer-service/internal/pkg/utils"
	"context"
)

type IAccountService interface {
	CreateUser(ctx context.Context, ae models.AccountEvent) error
}

type AccountService struct {
	producer *kafka.IProducer
	registry repository.Registry
}

func NewAccountService(p *kafka.IProducer, r repository.Registry) *AccountService {
	return &AccountService{
		producer: p,
		registry: r,
	}
}

func (as *AccountService) CreateUser(ctx context.Context, ae models.AccountEvent) error {
	user, err := as.shouldCreateUser(ctx, ae)
	if err != nil {
		return err
	}
	if user != nil {
		utils.Logger.Error("user already exists", err)
		return err
	} else {
		newUser := models.UserDBModel{
			Name:  ae.Name,
			Email: ae.Email,
		}

		userRepository := as.registry.UserRepository()

		err := userRepository.Create(ctx, newUser)
		if err != nil {
			utils.Logger.Error("error during create user", err)
			return err
		}

		user, err = as.shouldCreateUser(ctx, ae)
		if err != nil {
			return err
		}

		ae.Id = user.Id

		err = as.createAddress(ctx, ae)
		if err != nil {
			utils.Logger.Error("error during create address", err)
			return err
		}

		err = as.createPhone(ctx, ae)
		if err != nil {
			utils.Logger.Error("error during create phone", err)
			return err
		}

	}

	return nil
}

func (as *AccountService) shouldCreateUser(ctx context.Context, ae models.AccountEvent) (*models.UserDBModel, error) {
	userRequest := models.UserRequestByEmail{
		Email: ae.Email,
	}

	userRepository := as.registry.UserRepository()

	userDb, err := userRepository.GetByEmail(ctx, userRequest)
	if err != nil {
		utils.Logger.Error("error during get user by email", err)
		return nil, err
	}

	return userDb, nil
}

func (as *AccountService) createAddress(ctx context.Context, ae models.AccountEvent) error {
	address := models.AddressDBModel{
		Alias:       ae.Alias,
		City:        ae.City,
		District:    ae.District,
		PublicPlace: ae.PublicPlace,
		ZipCode:     ae.ZipCode,
		UserId:      ae.Id,
	}

	addressRepository := as.registry.AddressRepository()

	err := addressRepository.Create(ctx, address)
	if err != nil {
		utils.Logger.Error("error during create address", err)
		return err
	}

	return nil
}

func (as *AccountService) createPhone(ctx context.Context, ae models.AccountEvent) error {
	phone := models.PhoneDBModel{
		AreaCode:    ae.AreaCode,
		CountryCode: ae.CountryCode,
		Number:      ae.Number,
		UserId:      ae.Id,
	}

	phoneRepository := as.registry.PhoneRepository()

	err := phoneRepository.Create(ctx, phone)
	if err != nil {
		utils.Logger.Error("error during create phone", err)
		return err
	}

	return nil
}
