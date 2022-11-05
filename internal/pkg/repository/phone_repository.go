package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type PhoneRepository struct {
	scylla *db.Scylla
}

func NewPhoneRepository(s *db.Scylla) *PhoneRepository {
	return &PhoneRepository{
		scylla: s,
	}
}

func (repo *PhoneRepository) Insert(ctx context.Context, p models.Phone) *errorx.Error {
	stmt := `INSERT INTO phone (id, area_code, country_code, number) VALUES (uuid(),?,?,?)`
	err := repo.scylla.Insert(stmt, ctx, p.CountryCode, p.AreaCode, p.Number)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *PhoneRepository) GetById(ctx context.Context, p models.PhoneRequestById) (*models.Phone, *errorx.Error) {
	stmt := `SELECT id, area_code, country_code, number FROM phone WHERE id = ? LIMIT 1`
	rows := repo.scylla.GetById(stmt, ctx, p.Id)
	scan, err := repo.scanById(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *PhoneRepository) List(ctx context.Context) ([]models.Phone, *errorx.Error) {
	stmt := `SELECT * FROM phone`
	rows := repo.scylla.List(stmt, ctx)
	scan, err := repo.scanList(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *PhoneRepository) scanById(rows *gocql.Query) (*models.Phone, error) {
	p := models.Phone{}
	err := rows.Scan(
		&p.Id,
		&p.CountryCode,
		&p.AreaCode,
		&p.Number,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (repo *PhoneRepository) scanList(rows *gocql.Iter) ([]models.Phone, error) {
	pList := []models.Phone{}
	p := models.Phone{}
	scan := rows.Scanner()
	for scan.Next() {
		err := scan.Scan(&p.Id, &p.AreaCode, &p.CountryCode, &p.Number)
		if err != nil {
			return nil, err
		}
		pList = append(pList, p)
	}
	return pList, nil
}
