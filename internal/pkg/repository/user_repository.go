package repository

import (
	"account-consumer-service/internal/entities"
	"context"
	"strings"

	"github.com/gocql/gocql"
	"github.com/joomcode/errorx"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, u entities.User) *errorx.Error
	GetById(ctx context.Context, u entities.User) (*entities.User, *errorx.Error)
	List(ctx context.Context) ([]entities.User, *errorx.Error)
}

type UserRepository struct {
	conn *gocql.Session
}

func NewUserRepository(s *gocql.Session) *UserRepository {
	return &UserRepository{
		conn: s,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u entities.User) *errorx.Error {
	err := repo.conn.Query(`INSERT INTO user (id, address_id, phone_id, name, email) VALUES (uuid(),?,?,?,?)`,
		strings.ToLower(u.AddressId),
		strings.ToLower(u.PhoneId),
		strings.ToUpper(u.Name),
		strings.ToUpper(u.Email),
	).WithContext(ctx).Exec()
	if err != nil {
		return errorx.Decorate(err, "error during insert query")
	}
	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, u entities.User) (*entities.User, *errorx.Error) {
	user := entities.User{}
	err := repo.conn.Query(`SELECT id, address_id, phone_id, name, email FROM user WHERE id = ? LIMIT 1`,
		u.Id).WithContext(ctx).Consistency(gocql.One).Scan(
		&user.Id,
		&user.AddressId,
		&user.PhoneId,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}
	return &user, nil
}

func (repo *UserRepository) List(ctx context.Context) ([]entities.User, *errorx.Error) {
	scanner := repo.conn.Query(`SELECT id, address_id, phone_id, name, email FROM user`).WithContext(ctx).Iter().Scanner()
	uList := []entities.User{}
	u := entities.User{}
	for scanner.Next() {
		err := scanner.Scan(&u.Id, &u.AddressId, &u.PhoneId, &u.Name, &u.Email)
		if err != nil {
			return nil, errorx.Decorate(err, "error during scanner")
		}
		uList = append(uList, u)
	}
	if err := scanner.Err(); err != nil {
		return nil, errorx.Decorate(err, "error during select query")
	}
	return uList, nil
}
