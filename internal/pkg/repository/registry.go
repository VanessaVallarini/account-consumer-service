package repository

import (
	"account-consumer-service/internal/pkg/db"
)

type IRegistry interface {
	AddressRepository() IAddressRepository
	PhoneRepository() IPhoneRepository
	UserRepository() IUserRepository
}

type Registry struct {
	scylla db.IScylla
}

func NewRegistry(scylla db.IScylla) *Registry {
	return &Registry{scylla: scylla}
}

func (reg *Registry) AddressRepository() IAddressRepository {
	return NewAddressRepository(reg.scylla)
}

func (reg *Registry) PhoneRepository() IPhoneRepository {
	return NewPhoneRepository(reg.scylla)
}

func (reg *Registry) UserRepository() IUserRepository {
	return NewUserRepository(reg.scylla)
}
