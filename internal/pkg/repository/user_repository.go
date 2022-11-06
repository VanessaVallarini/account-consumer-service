package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"context"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type UserRepository struct {
	scylla db.ScyllaInterface
}

func NewUserRepository(s db.ScyllaInterface) *UserRepository {
	return &UserRepository{
		scylla: s,
	}
}

func (repo *UserRepository) Insert(ctx context.Context, u models.User) *errorx.Error {
	stmt := `INSERT INTO user (id, address_id, phone_id, email, name) VALUES (uuid(),?,?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, u.AddressId, u.PhoneId, u.Email, u.Name)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, u models.UserRequestById) (*models.User, *errorx.Error) {
	stmt := `SELECT id, address_id, phone_id, email, name FROM user WHERE id = ? LIMIT 1`
	rows := repo.scylla.GetById(ctx, stmt, u.Id)
	scan, err := repo.scanById(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *UserRepository) List(ctx context.Context) ([]models.User, *errorx.Error) {
	stmt := `SELECT * FROM user`
	rows := repo.scylla.List(ctx, stmt)
	scan, err := repo.scanList(rows)
	if err != nil {
		return nil, errorx.Decorate(err, "error during scan")
	}
	return scan, nil
}

func (repo *UserRepository) Update(ctx context.Context, u models.User) *errorx.Error {
	stmt := `UPDATE user SET address_id = ?, phone_id = ?, email = ?, name = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, u.AddressId, u.PhoneId, u.Email, u.Name, u.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *UserRepository) Delete(ctx context.Context, u models.UserRequestById) *errorx.Error {
	stmt := `DELETE from user WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, u.Id)
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
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
