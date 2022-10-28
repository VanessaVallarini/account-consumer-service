package repository

import "github.com/gocql/gocql"

type Registry interface {
	AddressRepository() AddressRepositoryInterface
	PhoneRepository() PhoneRepositoryInterface
	UserRepository() UserRepositoryInterface
}

type RegistryRepository struct {
	conn *gocql.Session
}

func NewRegistryRepository(s *gocql.Session) *RegistryRepository {
	return &RegistryRepository{
		conn: s,
	}
}

func (reg *RegistryRepository) AddressRepository() AddressRepositoryInterface {
	return NewAddressRepository(reg.conn)
}

func (reg *RegistryRepository) PhoneRepository() PhoneRepositoryInterface {
	return NewPhoneRepository(reg.conn)
}

func (reg *RegistryRepository) UserRepository() UserRepositoryInterface {
	return NewUserRepository(reg.conn)
}
