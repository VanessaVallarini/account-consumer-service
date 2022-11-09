package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/joomcode/errorx"
)

type PhoneRepositoryInterface interface {
	Insert(ctx context.Context, p models.Phone) *errorx.Error
	List(ctx context.Context) ([]models.Phone, *errorx.Error)
	Update(ctx context.Context, p models.Phone) *errorx.Error
	Delete(ctx context.Context, a models.PhoneRequestById) *errorx.Error
}

type PhoneRepository struct {
	scylla db.ScyllaInterface
}

func NewPhoneRepository(s db.ScyllaInterface) *PhoneRepository {
	return &PhoneRepository{
		scylla: s,
	}
}

func (repo *PhoneRepository) Insert(ctx context.Context, p models.Phone) *errorx.Error {
	stmt := `INSERT INTO phone (id, area_code, country_code, number) VALUES (uuid(),?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, p.CountryCode, p.AreaCode, p.Number)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

/* func (repo *PhoneRepository) GetById(ctx context.Context, p models.PhoneRequestById) (*models.Phone, *errorx.Error) {
	stmt := `SELECT id, area_code, country_code, number FROM phone WHERE id = ? LIMIT 1`
	rows := repo.scylla.GetById(ctx, stmt, p.Id)
	scan, err := repo.scanById(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
} */

/* func (repo *PhoneRepository) List(ctx context.Context) ([]models.Phone, *errorx.Error) {
	stmt := `SELECT * FROM phone`
	rows := repo.scylla.List(ctx, stmt)
	scan, err := repo.scanList(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
} */

func (repo *PhoneRepository) Update(ctx context.Context, p models.Phone) *errorx.Error {
	stmt := `UPDATE phone SET area_code = ?, country_code = ?, number = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, p.AreaCode, p.CountryCode, p.Number, p.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *PhoneRepository) Delete(ctx context.Context, u models.PhoneRequestById) *errorx.Error {
	stmt := `DELETE from phone WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, u.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}
