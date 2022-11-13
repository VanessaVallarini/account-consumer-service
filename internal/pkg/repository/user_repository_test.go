package repository

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/db/mocks"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRepository(t *testing.T) {
	scylla := mocks.NewScylla()
	userRepository := NewUserRepository(scylla)

	assert.NotNil(t, userRepository)
}

func TestCreateUser(t *testing.T) {
	t.Run("Expect to return success on create user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserDBModel{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := userRepository.Create(ctx, u)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserDBModel{}

		scylla.When("Insert",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query create user"),
		)

		err := userRepository.Create(ctx, u)

		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Expect to return success on update user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserDBModel{}

		scylla.When("Update",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := userRepository.Update(ctx, u)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserDBModel{}

		scylla.When("Update",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query update user"),
		)

		err := userRepository.Update(ctx, u)

		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Expect to return success on delete user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		p := models.UserRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		err := userRepository.Delete(ctx, p)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		p := models.UserRequestById{}

		scylla.When("Delete",
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query delete user"),
		)

		err := userRepository.Delete(ctx, p)

		assert.Error(t, err)
	})
}

func TestGetByIdUser(t *testing.T) {
	t.Run("Expect to return success on get user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		phone, err := userRepository.GetById(ctx, u)

		assert.Nil(t, err)
		assert.NotNil(t, phone)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserRequestById{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by id"),
		)

		user, err := userRepository.GetById(ctx, u)

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestGetByEmailUser(t *testing.T) {
	t.Run("Expect to return success on get user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserRequestByEmail{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			nil,
		)

		phone, err := userRepository.GetByEmail(ctx, u)

		assert.Nil(t, err)
		assert.NotNil(t, phone)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		u := models.UserRequestByEmail{}

		scylla.When("ScanMap",
			mock.Any,
			mock.Any,
			mock.Any,
			mock.Any,
		).Return(
			errors.New("error during query get user by email"),
		)

		user, err := userRepository.GetByEmail(ctx, u)

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("Expect to return success on get all user", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		randomUUID, _ := gocql.RandomUUID()
		var udList []models.UserDBModel
		ud := models.UserDBModel{
			Id:    randomUUID.String(),
			Name:  "Name",
			Email: "Email",
		}
		udList = append(udList, ud)
		ud = models.UserDBModel{
			Id:    randomUUID.String(),
			Name:  "Name",
			Email: "Email",
		}
		udList = append(udList, ud)

		var requestAsMap []map[string]interface{}
		marshalledRequest, _ := json.Marshal(udList)
		json.Unmarshal(marshalledRequest, &requestAsMap)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			requestAsMap,
			nil,
		)

		userList, err := userRepository.List(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, userList)
	})

	t.Run("Expect to return error during query", func(t *testing.T) {
		ctx := context.Background()
		scylla := mocks.NewScylla()
		userRepository := NewUserRepository(scylla)

		scylla.When("ScanMapSlice",
			mock.Any,
			mock.Any,
		).Return(
			nil,
			errors.New("error during query get all user"),
		)

		userList, err := userRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, userList)
	})
}
