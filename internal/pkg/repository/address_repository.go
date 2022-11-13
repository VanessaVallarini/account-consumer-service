package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"encoding/json"
)

type IAddressRepository interface {
	Create(ctx context.Context, a models.AddressDBModel) error
	GetById(ctx context.Context, a models.AddressRequestById) (*models.AddressDBModel, error)
	List(ctx context.Context) ([]models.AddressDBModel, error)
	Update(ctx context.Context, a models.AddressDBModel) error
	Delete(ctx context.Context, a models.AddressRequestById) error
}

type AddressRepository struct {
	scylla db.IScylla
}

func NewAddressRepository(s db.IScylla) *AddressRepository {
	return &AddressRepository{
		scylla: s,
	}
}

func (repo *AddressRepository) Create(ctx context.Context, a models.AddressDBModel) error {
	stmt := `INSERT INTO address (id, alias, city, district, public_place ,zip_code, user_id) VALUES (uuid(),?,?,?,?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode, a.UserId)
	if err != nil {
		utils.Logger.Info("error during query create address", err)
		return err
	}
	return nil
}

func (repo *AddressRepository) GetById(ctx context.Context, a models.AddressRequestById) (*models.AddressDBModel, error) {
	stmt := `SELECT id, alias, city, district, public_place, zip_code, user_id FROM address WHERE id = ? LIMIT 1`

	address := &models.AddressDBModel{}
	results := map[string]interface{}{
		"id":           &address.Id,
		"alias":        &address.Alias,
		"city":         &address.City,
		"district":     &address.District,
		"public_place": &address.PublicPlace,
		"zip_code":     &address.ZipCode,
		"user_id":      &address.UserId,
	}

	err := repo.scylla.ScanMap(ctx, stmt, results, a.Id)
	if err != nil {
		utils.Logger.Info("error during query get address by id", err)
		return nil, err
	}

	return address, nil
}

func (repo *AddressRepository) List(ctx context.Context) ([]models.AddressDBModel, error) {
	stmt := `SELECT * FROM address`

	aList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Info("error during query get all address", err)
		return nil, err
	}

	convertToAdressList := repo.scanAdressList(aList)

	return convertToAdressList, nil
}

func (repo *AddressRepository) Update(ctx context.Context, a models.AddressDBModel) error {
	stmt := `UPDATE address SET alias = ?, city = ?, district = ?, public_place = ?, zip_code = ?, user_id = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode, a.UserId, a.Id)
	if err != nil {
		utils.Logger.Info("error during query update address", err)
		return err
	}
	return nil
}

func (repo *AddressRepository) Delete(ctx context.Context, a models.AddressRequestById) error {
	stmt := `DELETE from address WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Id)
	if err != nil {
		utils.Logger.Info("error during query delete address", err)
		return err
	}
	return nil
}

func (repo *AddressRepository) scanAdressList(results []map[string]interface{}) []models.AddressDBModel {
	var aList []models.AddressDBModel

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &aList)

	return aList
}
