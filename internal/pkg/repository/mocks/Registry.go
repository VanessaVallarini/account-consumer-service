package mocks

import (
	"account-consumer-service/internal/pkg/repository"

	"github.com/maraino/go-mock"
)

type IRegistry interface {
	AccountRepository() IAccountRepository
}

type Registry struct {
	mock.Mock
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (m *Registry) AccountRepository() repository.IAccountRepository {
	return NewAccountRepository()
}
