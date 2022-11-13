package repository

import (
	"account-consumer-service/internal/pkg/db"
)

type IRegistry interface {
	AccountRepository() IAccountRepository
}

type Registry struct {
	scylla db.IScylla
}

func NewRegistry(scylla db.IScylla) *Registry {
	return &Registry{scylla: scylla}
}

func (reg *Registry) AccountRepository() IAccountRepository {
	return NewAccountRepository(reg.scylla)
}
