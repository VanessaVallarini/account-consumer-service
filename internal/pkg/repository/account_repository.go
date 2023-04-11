package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"strings"
)

type IAccountRepository interface {
	CreateOrUpdate(ctx context.Context, a models.Account) error
	Delete(ctx context.Context, a models.AccountRequestByEmail) error
	GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error)
}

type AccountRepository struct {
	scylla db.IScylla
}

func NewAccountRepository(s db.IScylla) *AccountRepository {
	return &AccountRepository{
		scylla: s,
	}
}

func (repo *AccountRepository) CreateOrUpdate(ctx context.Context, a models.Account) error {
	stmt := `INSERT INTO account 
				(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
			VALUES
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	err := repo.scylla.Insert(ctx, stmt, a.Email, a.FullNumber, a.Alias, a.City, a.DateTime, a.District, a.Name, a.PublicPlace, a.Status, a.ZipCode)
	if err != nil {
		utils.Logger.Error("account producer failed during query create account: %v", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) Delete(ctx context.Context, a models.AccountRequestByEmail) error {
	stmt := `DELETE from account WHERE email = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Email)
	if err != nil {
		utils.Logger.Error("account producer failed during query delete account: %v", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	stmt := `SELECT * FROM account WHERE email = ?`
	account := &models.Account{}
	results := map[string]interface{}{
		"email":        &account.Email,
		"full_number":  &account.FullNumber,
		"alias":        &account.Alias,
		"city":         &account.City,
		"date_time":    &account.DateTime,
		"district":     &account.District,
		"name":         &account.Name,
		"public_place": &account.PublicPlace,
		"status":       &account.Status,
		"zip_code":     &account.ZipCode,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Error("account producer failed during query get account by email: %v", err)
		return nil, err
	}

	return account, nil
}
