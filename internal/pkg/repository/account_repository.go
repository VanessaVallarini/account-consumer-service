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
	Create(ctx context.Context, a models.AccountCreate) error
	GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error)
	GetByFullNumber(ctx context.Context, a models.AccountRequestByFullNumber) (*models.Account, error)
	GetBy(ctx context.Context, a models.AccountRequestBy) (*models.Account, error)
	List(ctx context.Context) ([]models.Account, error)
	Update(ctx context.Context, a models.Account) error
	Delete(ctx context.Context, a models.AccountRequestById) error
}

type AccountRepository struct {
	scylla db.IScylla
}

func NewAccountRepository(s db.IScylla) *AccountRepository {
	return &AccountRepository{
		scylla: s,
	}
}

func (repo *AccountRepository) Create(ctx context.Context, a models.AccountCreate) error {
	stmt := `INSERT INTO account 
				(id, alias, city, district, email, full_number, name, public_place, zip_code)
			VALUES
				(uuid(), ?, ?, ?, ?, ?, ?, ?, ?);`
	err := repo.scylla.Insert(ctx, stmt, a.Alias, a.City, a.District, a.Email, a.FullNumber, a.Name, a.PublicPlace, a.ZipCode)
	if err != nil {
		utils.Logger.Error("error during query create account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	stmt := `SELECT * FROM account WHERE email = ? and full_number = ? LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"email":        &account.Email,
		"full_number":  &account.FullNumber,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"name":         &account.Name,
		"public_place": &account.PublicPlace,
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

func (repo *AccountRepository) GetByFullNumber(ctx context.Context, a models.AccountRequestByFullNumber) (*models.Account, error) {
	stmt := `SELECT * FROM account WHERE full_number = ? LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"email":        &account.Email,
		"full_number":  &account.FullNumber,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"name":         &account.Name,
		"public_place": &account.PublicPlace,
		"zip_code":     &account.ZipCode,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.FullNumber)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Error("error during query get account by full number", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetBy(ctx context.Context, a models.AccountRequestBy) (*models.Account, error) {
	stmt := `SELECT * FROM account WHERE email = ? LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"email":        &account.Email,
		"full_number":  &account.FullNumber,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"name":         &account.Name,
		"public_place": &account.PublicPlace,
		"zip_code":     &account.ZipCode,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Error("error during query get account by", err)
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

func (repo *AccountRepository) Update(ctx context.Context, a models.Account) error {
	stmt := `UPDATE account SET 
			 alias = ?, city = ?, district = ?, email = ?, full_number = ?, name = ?, public_place = ?, zip_code = ?
		     WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, a.Alias, a.City, a.District, a.Email, a.FullNumber, a.Name, a.PublicPlace, a.ZipCode, a.Id)
	if err != nil {
		utils.Logger.Error("error during query update account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) Delete(ctx context.Context, a models.AccountRequestById) error {
	stmt := `DELETE from account WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Id)
	if err != nil {
		utils.Logger.Error("error during query delete account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) scanAccountList(results []map[string]interface{}) []models.Account {
	var aList []models.Account

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &aList)

	return aList
}
