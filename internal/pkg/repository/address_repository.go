package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/scylla"
	"context"

	"github.com/joomcode/errorx"
)

type AddressRepository struct {
	scylla *scylla.IScylla
}

func NewAddressRepository(s *scylla.IScylla) *AddressRepository {
	return &AddressRepository{
		scylla: s,
	}
}

func (repo *AddressRepository) Insert(ctx context.Context, a models.Address) *errorx.Error {
	stmt := `INSERT INTO address (id, alias, city, district, public_place ,zip_code) VALUES (uuid(),?,?,?,?,?)`
	err := repo.scylla.Insert(stmt, ctx, a)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *AddressRepository) GetById(ctx context.Context, a models.AddressRequestById) (*models.Address, *errorx.Error) {
	stmt := `SELECT id, alias, city, district, public_place, zip_code FROM address WHERE id = ? LIMIT 1`
	exec := repo.scylla.GetById(stmt, ctx, a)
	address := models.Address{}
	scan := exec.Scan(
		&address.Id,
		&address.Alias,
		&address.City,
		&address.District,
		&address.PublicPlace,
		&address.ZipCode,
	)
	if scan != nil {
		return nil, errorx.Decorate(scan, "error during scan")
	}
	return &address, nil
}

func (repo *AddressRepository) List(ctx context.Context) ([]models.Address, *errorx.Error) {
	stmt := `SELECT id, alias, city, district, public_place, zip_code FROM address WHERE id = ? LIMIT 1`
	exec := repo.scylla.List(stmt, ctx)
	aList := []models.Address{}
	a := models.Address{}
	scan := exec.Scanner()
	for scan.Next() {
		err := scan.Scan(&a.Id, &a.Alias, &a.City, &a.District, &a.PublicPlace, &a.ZipCode)
		if err != nil {
			return nil, errorx.Decorate(err, "error during scanner")
		}
		aList = append(aList, a)
	}
	return aList, nil
}
