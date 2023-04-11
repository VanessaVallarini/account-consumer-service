package db

import (
	"account-consumer-service/internal/models"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScylla(t *testing.T) {
	configDatabase := models.DatabaseConfig{
		DatabaseUser:     "cassandra",
		DatabasePassword: "cassandra",
		DatabaseKeyspace: "account_consumer_service",

		DatabaseHost:                "localhost",
		DatabasePort:                9042,
		DatabaseConnectionRetryTime: 5,
		DatabaseRetryMinArg:         1,
		DatabaseRetryMaxArg:         10,
		DatabaseNumRetries:          5,
		DatabaseClusterTimeout:      5,
	}

	scylla, _ := NewScylla(&configDatabase)

	assert.NotNil(t, scylla)
}

func TestScanMap(t *testing.T) {
	t.Run("Expect to return success on get account by email", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		stmt = `SELECT * FROM account WHERE email = ?`
		account := &models.Account{}
		results := map[string]interface{}{
			"email":        &account.Email,
			"full_number":  &account.FullNumber,
			"alias":        &account.Alias,
			"city":         &account.City,
			"date_time":    &account.DateTime,
			"district":     &account.District,
			"name":         &account.Name,
			"public_place": &account.PublicPlace,
			"status":       &account.Status,
			"zip_code":     &account.ZipCode,
		}

		err := scylla.ScanMap(ctx, stmt, results, accountCreate.Email)

		assert.Nil(t, err)
		assert.NotNil(t, account)
	})

	t.Run("Expect to return error during query on get account by email when account does not exist", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		stmt := `SELECT * FROM account WHERE email = ?`
		account := &models.Account{}
		results := map[string]interface{}{
			"email":        &account.Email,
			"full_number":  &account.FullNumber,
			"alias":        &account.Alias,
			"city":         &account.City,
			"date_time":    &account.DateTime,
			"district":     &account.District,
			"name":         &account.Name,
			"public_place": &account.PublicPlace,
			"status":       &account.Status,
			"zip_code":     &account.ZipCode,
		}

		err := scylla.ScanMap(ctx, stmt, results, "teste")

		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "not found"))
	})

	t.Run("Expect to return error during query on get account by email when stm is invalid", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		stmt := `SELECT FROM account WHERE email = ?`
		account := &models.Account{}
		results := map[string]interface{}{
			"email":        &account.Email,
			"full_number":  &account.FullNumber,
			"alias":        &account.Alias,
			"city":         &account.City,
			"date_time":    &account.DateTime,
			"district":     &account.District,
			"name":         &account.Name,
			"public_place": &account.PublicPlace,
			"status":       &account.Status,
			"zip_code":     &account.ZipCode,
		}

		err := scylla.ScanMap(ctx, stmt, results, "teste")

		assert.Error(t, err)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Expect to return success on create or update account", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		err := scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query create or update account when arguments is invalid", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		err := scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status)

		assert.Error(t, err)
	})

	t.Run("Expect to return error during query create or update account when stm is invalid", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		err := scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Expect to return success on delete account", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "loremDelete@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		stmt = `DELETE from account WHERE email = ?`

		err := scylla.Delete(ctx, stmt, accountCreate.Email)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query delete account when arguments is invalid", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "loremDelete@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		stmt = `DELETE from account WHERE email = ?`

		err := scylla.Delete(ctx, stmt)

		assert.Error(t, err)
	})

	t.Run("Expect to return error during query delete account when stm is invalid", func(t *testing.T) {
		configDatabase := models.DatabaseConfig{
			DatabaseUser:     "cassandra",
			DatabasePassword: "cassandra",
			DatabaseKeyspace: "account_consumer_service",

			DatabaseHost:                "localhost",
			DatabasePort:                9042,
			DatabaseConnectionRetryTime: 5,
			DatabaseRetryMinArg:         1,
			DatabaseRetryMaxArg:         10,
			DatabaseNumRetries:          5,
			DatabaseClusterTimeout:      5,
		}
		scylla, _ := NewScylla(&configDatabase)
		ctx := context.Background()

		accountCreate := &models.Account{
			Email:       "loremDelete@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}
		stmt := `INSERT INTO account 
					(email, full_number, alias, city, date_time, district, name, public_place, status, zip_code)
				VALUES
					(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
		scylla.Insert(ctx, stmt, accountCreate.Email, accountCreate.FullNumber, accountCreate.Alias, accountCreate.City,
			accountCreate.DateTime, accountCreate.District, accountCreate.Name, accountCreate.PublicPlace, accountCreate.Status,
			accountCreate.ZipCode)

		stmt = `DELETE account WHERE email = ?`

		err := scylla.Delete(ctx, stmt, accountCreate.Email)

		assert.Error(t, err)
	})
}
