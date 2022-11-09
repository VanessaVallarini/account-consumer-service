package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/joomcode/errorx"
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

	return address, nil
}

func (repo *AddressRepository) List(ctx context.Context) ([]map[string]interface{}, *errorx.Error) {
	stmt := `SELECT * FROM address`

	aList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		return nil, errorx.Decorate(err, "error during query")
	}

	return aList, nil
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
