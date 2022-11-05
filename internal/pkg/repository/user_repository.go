package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type UserRepository struct {
	scylla *db.Scylla
}

func NewUserRepository(s *db.Scylla) *UserRepository {
	return &UserRepository{
		scylla: s,
	}
}

func (repo *UserRepository) Insert(ctx context.Context, u models.User) *errorx.Error {
	stmt := `INSERT INTO user (id, address_id, phone_id, email, name) VALUES (uuid(),?,?,?,?)`
	err := repo.scylla.Insert(stmt, ctx, u.AddressId, u.PhoneId, u.Email, u.Name)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, u models.UserRequestById) (*models.User, *errorx.Error) {
	stmt := `SELECT id, address_id, phone_id, email, name FROM user WHERE id = ? LIMIT 1`
	rows := repo.scylla.GetById(stmt, ctx, u.Id)
	scan, err := repo.scanById(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *UserRepository) List(ctx context.Context) ([]models.User, *errorx.Error) {
	stmt := `SELECT * FROM user`
	rows := repo.scylla.List(stmt, ctx)
	scan, err := repo.scanList(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *UserRepository) scanById(rows *gocql.Query) (*models.User, error) {
	u := models.User{}
	err := rows.Scan(
		&u.Id,
		&u.AddressId,
		&u.PhoneId,
		&u.Email,
		&u.Name,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (repo *UserRepository) scanList(rows *gocql.Iter) ([]models.User, error) {
	uList := []models.User{}
	u := models.User{}
	scan := rows.Scanner()
	for scan.Next() {
		err := scan.Scan(&u.Id, &u.AddressId, &u.PhoneId, &u.Email, &u.Name)
		if err != nil {
			return nil, err
		}
		uList = append(uList, u)
	}
	return uList, nil
}
