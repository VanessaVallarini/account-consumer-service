package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"encoding/json"
	"strings"
)

type IAccountRepository interface {
	Create(ctx context.Context, a models.Account) error
	Update(ctx context.Context, a models.Account) error
	Delete(ctx context.Context, a models.AccountRequestByEmailAndFullNumber) error
	GetByEmailAndFullNumber(ctx context.Context, a models.AccountRequestByEmailAndFullNumber) (*models.Account, error)
	GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error)
	GetByFullNumber(ctx context.Context, a models.AccountRequestByFullNumber) (*models.Account, error)
	List(ctx context.Context) ([]models.Account, error)
}

type AccountRepository struct {
	scylla db.IScylla
}

func NewAccountRepository(s db.IScylla) *AccountRepository {
	return &AccountRepository{
		scylla: s,
	}
}

func (repo *AccountRepository) Create(ctx context.Context, a models.Account) error {
	stmt := `INSERT INTO account 
				(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
			VALUES
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	err := repo.scylla.Insert(ctx, stmt, a.Email, a.FullNumber, a.Alias, a.City, a.DateTime, a.District, a.Name, a.PublicPlace, a.Status, a.ZipCode)
	if err != nil {
		utils.Logger.Error("error during query create account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) Update(ctx context.Context, a models.Account) error {
	stmt := `INSERT INTO account 
				(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
			VALUES
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	err := repo.scylla.Insert(ctx, stmt, a.Email, a.FullNumber, a.Alias, a.City, a.DateTime, a.District, a.Name, a.PublicPlace, a.Status, a.ZipCode)
	if err != nil {
		utils.Logger.Error("error during query create account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) Delete(ctx context.Context, a models.AccountRequestByEmailAndFullNumber) error {
	stmt := `DELETE from account WHERE email = ? AND full_number = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Email, a.FullNumber)
	if err != nil {
		utils.Logger.Error("error during query delete account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) GetByEmailAndFullNumber(ctx context.Context, a models.AccountRequestByEmailAndFullNumber) (*models.Account, error) {
	stmt := `SELECT email, full_number, alias, city, CAST(date_time AS text) , district, name, public_place, status, zip_code FROM account WHERE email = ? and full_number = ? LIMIT 1`
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
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Email, a.FullNumber)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Error("error during query get account by email", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	stmt := `SELECT email, full_number, alias, city, CAST(date_time AS text) , district, name, public_place, status, zip_code FROM account WHERE email = ? LIMIT 1`
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
		utils.Logger.Error("error during query get account by email", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetByFullNumber(ctx context.Context, a models.AccountRequestByFullNumber) (*models.Account, error) {
	stmt := `SELECT email, full_number, alias, city, CAST(date_time AS text) , district, name, public_place, status, zip_code FROM account WHERE full_number = ? LIMIT 1`
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
	err := repo.scylla.ScanMap(ctx, stmt, results, a.FullNumber)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Error("error during query get account by email", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) List(ctx context.Context) ([]models.Account, error) {
	stmt := `SELECT * FROM account`

	uList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Error("error during query get all account", err)
		return nil, err
	}

	convertToUserList := repo.scanAccountList(uList)

	return convertToUserList, nil
}

func (repo *AccountRepository) scanAccountList(results []map[string]interface{}) []models.Account {
	var aList []models.Account

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &aList)

	return aList
}
