package repository

import (
	"account-consumer-service/internal/entities"
	"context"
	"strings"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type AddressRepositoryInterface interface {
	Create(ctx context.Context, a entities.Address) *errorx.Error
	GetById(ctx context.Context, address entities.Address) (*entities.Address, *errorx.Error)
	List(ctx context.Context) ([]entities.Address, *errorx.Error)
}

type AddressRepository struct {
	conn *gocql.Session
}

func NewAddressRepository(s *gocql.Session) *AddressRepository {
	return &AddressRepository{
		conn: s,
	}
}

func (repo *AddressRepository) Create(ctx context.Context, a entities.Address) *errorx.Error {
	err := repo.conn.Query(`INSERT INTO address (id, alias, city, district, public_place ,zip_code) VALUES (uuid(),?,?,?,?,?)`,
		strings.ToUpper(a.Alias),
		strings.ToUpper(a.City),
		strings.ToUpper(a.District),
		strings.ToUpper(a.PublicPlace),
		strings.ToUpper(a.ZipCode),
	).WithContext(ctx).Exec()
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *AddressRepository) GetById(ctx context.Context, address entities.Address) (*entities.Address, *errorx.Error) {
	a := entities.Address{}
	err := repo.conn.Query(`SELECT id, alias, city, district, public_place, zip_code FROM address WHERE id = ? LIMIT 1`,
		address.Id).WithContext(ctx).Consistency(gocql.One).Scan(
		&a.Id,
		&a.Alias,
		&a.City,
		&a.District,
		&a.PublicPlace,
		&a.ZipCode,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}
	return &a, nil
}

func (repo *AddressRepository) List(ctx context.Context) ([]entities.Address, *errorx.Error) {
	scanner := repo.conn.Query(`SELECT id, alias, city, district, public_place, zip_code FROM address`).WithContext(ctx).Iter().Scanner()
	aList := []entities.Address{}
	a := entities.Address{}
	for scanner.Next() {
		err := scanner.Scan(&a.Id, &a.Alias, &a.City, &a.District, &a.PublicPlace, &a.ZipCode)
		if err != nil {
			return nil, errorx.Decorate(err, "error during scanner")
		}
		aList = append(aList, a)
	}
	if err := scanner.Err(); err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}
	return aList, nil
}
