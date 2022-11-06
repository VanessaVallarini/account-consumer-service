package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
	"github.com/mitchellh/mapstructure"
)

type FormRepositoryInterface interface {
	Insert(ctx context.Context, a models.Address) *errorx.Error
	GetById(ctx context.Context, a models.AddressRequestById) (*models.Address, *errorx.Error)
	List(ctx context.Context) ([]models.Address, *errorx.Error)
	Update(ctx context.Context, a models.Address) *errorx.Error
	Delete(ctx context.Context, a models.AddressRequestById) *errorx.Error
}

type AddressRepository struct {
	scylla db.ScyllaInterface
}

func NewAddressRepository(s db.ScyllaInterface) *AddressRepository {
	return &AddressRepository{
		scylla: s,
	}
}

func (repo *AddressRepository) Insert(ctx context.Context, a models.Address) *errorx.Error {
	stmt := `INSERT INTO address (id, alias, city, district, public_place ,zip_code) VALUES (uuid(),?,?,?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *AddressRepository) GetById(ctx context.Context, a models.AddressRequestById) (*models.Address, *errorx.Error) {
	stmt := `SELECT id, alias, city, district, public_place, zip_code FROM address WHERE id = ? LIMIT 1`

	address := &models.Address{}
	results := map[string]interface{}{
		"id":           &address.Id,
		"alias":        &address.Alias,
		"city":         &address.City,
		"district":     &address.District,
		"public_place": &address.PublicPlace,
		"zip_code":     &address.ZipCode,
	}

	err := repo.scylla.ScanMap(ctx, stmt, results, a.Id)
	if err != nil {
		return nil, errorx.Decorate(err, "error during query")
	}

	address, err = repo.convertToAdress(results)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}

	return address, nil
}

func (repo *AddressRepository) List(ctx context.Context) ([]models.Address, *errorx.Error) {
	stmt := `SELECT * FROM address`
	rows := repo.scylla.List(ctx, stmt)
	scan, err := repo.scanList(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *AddressRepository) Update(ctx context.Context, a models.Address) *errorx.Error {
	stmt := `UPDATE address SET alias = ?, city = ?, district = ?, public_place = ?, zip_code = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, a.Alias, a.City, a.District, a.PublicPlace, a.ZipCode, a.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *AddressRepository) Delete(ctx context.Context, a models.AddressRequestById) *errorx.Error {
	stmt := `DELETE from address WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *AddressRepository) convertToAdress(results map[string]interface{}) (*models.Address, error) {
	var a models.Address
	err := mapstructure.Decode(results, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (repo *AddressRepository) scanList(rows *gocql.Iter) ([]models.Address, error) {
	aList := []models.Address{}
	a := models.Address{}
	scan := rows.Scanner()
	for scan.Next() {
		err := scan.Scan(&a.Id, &a.Alias, &a.City, &a.District, &a.PublicPlace, &a.ZipCode)
		if err != nil {
			return nil, err
		}
		aList = append(aList, a)
	}
	return aList, nil
}
