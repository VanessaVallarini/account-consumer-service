package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"encoding/json"
)

type IUserRepository interface {
	Create(ctx context.Context, a models.UserDBModel) error
	GetById(ctx context.Context, a models.UserRequestById) (*models.UserDBModel, error)
	GetByEmail(ctx context.Context, u models.UserRequestByEmail) (*models.UserDBModel, error)
	List(ctx context.Context) ([]models.UserDBModel, error)
	Update(ctx context.Context, a models.UserDBModel) error
	Delete(ctx context.Context, a models.UserRequestById) error
}

type UserRepository struct {
	scylla db.IScylla
}

func NewUserRepository(s db.IScylla) *UserRepository {
	return &UserRepository{
		scylla: s,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u models.UserDBModel) error {
	stmt := `INSERT INTO user (id, email, name) VALUES (uuid(),?,?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, u.Email, u.Name)
	if err != nil {
		utils.Logger.Info("error during query create user", err)
		return err
	}
	return nil
}

func (repo *UserRepository) GetById(ctx context.Context, u models.UserRequestById) (*models.UserDBModel, error) {
	stmt := `SELECT id, email, name FROM user WHERE id = ? LIMIT 1`
	user := &models.UserDBModel{}
	results := map[string]interface{}{
		"id":    &user.Id,
		"email": &user.Email,
		"name":  &user.Name,
	}

	err := repo.scylla.ScanMap(ctx, stmt, results, u.Id)
	if err != nil {
		utils.Logger.Info("error during query get user by id", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, u models.UserRequestByEmail) (*models.UserDBModel, error) {
	stmt := `SELECT id, email, name FROM user WHERE email = ? LIMIT 1`
	user := &models.UserDBModel{}
	results := map[string]interface{}{
		"id":    &user.Id,
		"email": &user.Email,
		"name":  &user.Name,
	}

	err := repo.scylla.ScanMap(ctx, stmt, results, u.Email)
	if err != nil {
		utils.Logger.Info("error during query get user by email", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) List(ctx context.Context) ([]models.UserDBModel, error) {
	stmt := `SELECT * FROM user`

	uList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Info("error during query get all user", err)
		return nil, err
	}

	convertToUserList := repo.scanUserList(uList)

	return convertToUserList, nil
}

func (repo *UserRepository) Update(ctx context.Context, u models.UserDBModel) error {
	stmt := `UPDATE user SET email = ?, name = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, u.Email, u.Name, u.Id)
	if err != nil {
		utils.Logger.Info("error during query update user", err)
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(ctx context.Context, u models.UserRequestById) error {
	stmt := `DELETE from user WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, u.Id)
	if err != nil {
		utils.Logger.Info("error during query delete user", err)
		return err
	}
	return nil
}

func (repo *UserRepository) scanUserList(results []map[string]interface{}) []models.UserDBModel {
	var uList []models.UserDBModel

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &uList)

	return uList
}
