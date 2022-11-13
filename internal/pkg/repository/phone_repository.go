package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db"
	"account-consumer-service/internal/pkg/utils"
	"context"
	"encoding/json"
)

type IPhoneRepository interface {
	Create(ctx context.Context, a models.PhoneDBModel) error
	GetById(ctx context.Context, a models.PhoneRequestById) (*models.PhoneDBModel, error)
	List(ctx context.Context) ([]models.PhoneDBModel, error)
	Update(ctx context.Context, a models.PhoneDBModel) error
	Delete(ctx context.Context, a models.PhoneRequestById) error
}

type PhoneRepository struct {
	scylla db.IScylla
}

func NewPhoneRepository(s db.IScylla) *PhoneRepository {
	return &PhoneRepository{
		scylla: s,
	}
}

func (repo *PhoneRepository) Create(ctx context.Context, p models.PhoneDBModel) error {
	stmt := `INSERT INTO phone (id, area_code, country_code, number, user_id) VALUES (uuid(),?,?,?,?)`
	err := repo.scylla.Insert(ctx, stmt, p.CountryCode, p.AreaCode, p.Number, p.UserId)
	if err != nil {
		utils.Logger.Info("error during query create phone", err)
		return err
	}
	return nil
}

func (repo *PhoneRepository) GetById(ctx context.Context, p models.PhoneRequestById) (*models.PhoneDBModel, error) {
	stmt := `SELECT id, area_code, country_code, number FROM phone WHERE id = ? LIMIT 1`

	phone := &models.PhoneDBModel{}
	results := map[string]interface{}{
		"id":           &phone.Id,
		"area_code":    &phone.AreaCode,
		"country_code": &phone.CountryCode,
		"number":       &phone.Number,
		"user_id":      &phone.UserId,
	}

	err := repo.scylla.ScanMap(ctx, stmt, results, p.Id)
	if err != nil {
		utils.Logger.Info("error during query get phone by id", err)
		return nil, err
	}

	return phone, nil
}

func (repo *PhoneRepository) List(ctx context.Context) ([]models.PhoneDBModel, error) {
	stmt := `SELECT * FROM phone`

	pList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Info("error during query get all phone", err)
		return nil, err
	}

	convertToPhoneList := repo.scanPhoneList(pList)

	return convertToPhoneList, nil
}

func (repo *PhoneRepository) Update(ctx context.Context, p models.PhoneDBModel) error {
	stmt := `UPDATE phone SET area_code = ?, country_code = ?, number = ?, user_id = ? WHERE id = ?`
	err := repo.scylla.Update(ctx, stmt, p.AreaCode, p.CountryCode, p.Number, p.UserId, p.Id)
	if err != nil {
		utils.Logger.Info("error during query update phone", err)
		return err
	}
	return nil
}

func (repo *PhoneRepository) Delete(ctx context.Context, a models.PhoneRequestById) error {
	stmt := `DELETE from phone WHERE id = ?`
	err := repo.scylla.Delete(ctx, stmt, a.Id)
	if err != nil {
		utils.Logger.Info("error during query delete phone", err)
		return err
	}
	return nil
}

func (repo *PhoneRepository) scanPhoneList(results []map[string]interface{}) []models.PhoneDBModel {
	var pList []models.PhoneDBModel

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &pList)

	return pList
}
