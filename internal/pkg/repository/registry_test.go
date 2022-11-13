package repository

import (
	"account-consumer-service/internal/pkg/db/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRegistry(t *testing.T) *Registry {
	scylla := mocks.NewScylla()
	return NewRegistry(scylla)
}

func TestNewRegistry(t *testing.T) {
	reg := setupRegistry(t)

	assert.NotNil(t, reg)
}

func TestAccountRepositoryFromRegistry(t *testing.T) {
	reg := setupRegistry(t)

	AccountRepository := reg.AccountRepository()
	assert.NotNil(t, AccountRepository)
}
