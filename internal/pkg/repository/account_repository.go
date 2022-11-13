package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"encoding/json"
)

type IAccountRepository interface {
	Create(ctx context.Context, a models.Account) error
	GetById(ctx context.Context, a models.AccountRequestById) (*models.Account, error)
	GetByEmail(ctx context.Context, u models.AccountRequestByEmail) (*models.Account, error)
	GetByPhone(ctx context.Context, u models.AccountRequestByPhone) (*models.Account, error)
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

func (repo *AccountRepository) Create(ctx context.Context, a models.Account) error {
	stmt := `INSERT INTO account_consumer_service.account 
				(id,alias,area_code,city,country_code,district,email,name,number,public_place,zip_code)
			VALUES
				(uuid(),?,?,?,?,?,?,?,?,?,?);`
	err := repo.scylla.Insert(ctx, stmt, a.Alias, a.AreaCode, a.City, a.CountryCode, a.District, a.Email, a.Name, a.Number, a.PublicPlace, a.ZipCode)
	if err != nil {
		utils.Logger.Info("error during query create account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) GetById(ctx context.Context, a models.AccountRequestById) (*models.Account, error) {
	stmt := `SELECT id,alias,area_code,city,country_code,district,email,name,number,public_place,zip_code 
			 FROM account 
			 WHERE id = ? 
			 LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"name":         &account.Name,
		"email":        &account.Email,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"public_place": &account.PublicPlace,
		"zip_code":     &account.ZipCode,
		"country_code": &account.CountryCode,
		"area_code":    &account.AreaCode,
		"number":       &account.Number,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Id)
	if err != nil {
		utils.Logger.Info("error during query get account by id", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	stmt := `SELECT id,alias,area_code,city,country_code,district,email,name,number,public_place,zip_code 
	         FROM account 
			 WHERE email = ? 
			 LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"name":         &account.Name,
		"email":        &account.Email,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"public_place": &account.PublicPlace,
		"zip_code":     &account.ZipCode,
		"country_code": &account.CountryCode,
		"area_code":    &account.AreaCode,
		"number":       &account.Number,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Email)
	if err != nil {
		utils.Logger.Info("error during query get account by email", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) GetByPhone(ctx context.Context, a models.AccountRequestByPhone) (*models.Account, error) {
	stmt := `SELECT id,alias,area_code,city,country_code,district,email,name,number,public_place,zip_code 
	         FROM account
			 WHERE area_code = ?
			 AND  country_code = ?
			 AND number = ?
			 LIMIT 1`
	account := &models.Account{}
	results := map[string]interface{}{
		"id":           &account.Id,
		"name":         &account.Name,
		"email":        &account.Email,
		"alias":        &account.Alias,
		"city":         &account.City,
		"district":     &account.City,
		"public_place": &account.PublicPlace,
		"zip_code":     &account.ZipCode,
		"country_code": &account.CountryCode,
		"area_code":    &account.AreaCode,
		"number":       &account.Number,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.AreaCode, a.CountryCode, a.Number)
	if err != nil {
		utils.Logger.Info("error during query get account by phone", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) List(ctx context.Context) ([]models.Account, error) {
	stmt := `SELECT * FROM account`

	uList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Info("error during query get all account", err)
		return nil, err
	}

	convertToUserList := repo.scanAccountList(uList)

	return convertToUserList, nil
}

func (repo *AccountRepository) Update(ctx context.Context, a models.Account) error {
	stmt := `UPDATE user SET 
			 alias=?,area_code=?,city=?,country_code=?,district=?,email=?,name=?,number=?,public_place=?,zip_code=?
		     WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, a.Alias, a.AreaCode, a.City, a.CountryCode, a.District, a.Email, a.Name, a.Number, a.PublicPlace, a.ZipCode, a.Id)
	if err != nil {
		utils.Logger.Info("error during query update account", err)
		return err
	}
	return nil
}

func (repo *AccountRepository) Delete(ctx context.Context, a models.AccountRequestById) error {
	stmt := `DELETE from account WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Id)
	if err != nil {
		utils.Logger.Info("error during query delete account", err)
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
