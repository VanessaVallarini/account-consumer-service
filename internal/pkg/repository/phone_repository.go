package repository

import (
	"account-consumer-service/internal/entities"
	"context"
	"strings"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type PhoneRepositoryInterface interface {
	Create(ctx context.Context, p entities.Phone) *errorx.Error
	GetById(ctx context.Context, p entities.PhoneRequestById) (*entities.Phone, *errorx.Error)
	List(ctx context.Context) ([]entities.Phone, *errorx.Error)
}

type PhoneRepository struct {
	conn *gocql.Session
}

func NewPhoneRepository(s *gocql.Session) *PhoneRepository {
	return &PhoneRepository{
		conn: s,
	}
}

func (repo *PhoneRepository) Create(ctx context.Context, p entities.Phone) *errorx.Error {
	err := repo.conn.Query(`INSERT INTO phone (id, country_code, area_code, number) VALUES (uuid(),?,?,?)`,
		strings.ToUpper(p.CountryCode),
		strings.ToUpper(p.AreaCode),
		strings.ToUpper(p.Number),
	).WithContext(ctx).Exec()
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *PhoneRepository) GetById(ctx context.Context, p entities.PhoneRequestById) (*entities.Phone, *errorx.Error) {
	phone := entities.Phone{}
	err := repo.conn.Query(`SELECT id, country_code, area_code, number FROM phone WHERE id = ? LIMIT 1`,
		p.Id).WithContext(ctx).Consistency(gocql.One).Scan(
		&phone.Id,
		&phone.CountryCode,
		&phone.AreaCode,
		&phone.Number,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}

	return &phone, nil
}

func (repo *PhoneRepository) List(ctx context.Context) ([]entities.Phone, *errorx.Error) {
	scanner := repo.conn.Query(`SELECT id, country_code, area_code, number FROM phone`).WithContext(ctx).Iter().Scanner()
	pList := []entities.Phone{}
	p := entities.Phone{}
	for scanner.Next() {
		err := scanner.Scan(&p.Id, &p.CountryCode, &p.AreaCode, &p.Number)
		if err != nil {
			return nil, errorx.Decorate(err, "error during scanner")
		}
		pList = append(pList, p)
	}
	if err := scanner.Err(); err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}
	return pList, nil
}
