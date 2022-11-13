package mocks

import (
	"account-consumer-service/internal/models"
	"context"

	"github.com/maraino/go-mock"
)

type IAccountRepository interface {
	Create(ctx context.Context, a models.Account) error
	GetById(ctx context.Context, a models.AccountRequestById) (*models.Account, error)
	GetByEmail(ctx context.Context, u models.AccountRequestByEmail) (*models.Account, error)
	GetByPhone(ctx context.Context, u models.AccountRequestByPhone) (*models.Account, error)
	List(ctx context.Context) ([]models.Account, error)
	Update(ctx context.Context, a models.Account) error
	Delete(ctx context.Context, a models.AccountRequestById) error
}

// AccountRepository is an autogenerated mock type for the AccountRepositoryInterface type
type AccountRepository struct {
	mock.Mock
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

// Create provides a mock function with given fields: ctx, params
func (m *AccountRepository) Create(ctx context.Context, params models.Account) error {
	ret := m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Account) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, params
func (m *AccountRepository) GetById(ctx context.Context, params models.AccountRequestById) (*models.Account, error) {
	ret := m.Called(ctx, params)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(context.Context, models.AccountRequestById) *models.Account); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.AccountRequestById) error); ok {
		r1 = rf(ctx, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, params
func (m *AccountRepository) GetByEmail(ctx context.Context, params models.AccountRequestByEmail) (*models.Account, error) {
	ret := m.Called(ctx, params)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(context.Context, models.AccountRequestByEmail) *models.Account); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.AccountRequestByEmail) error); ok {
		r1 = rf(ctx, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

// GetByPhone provides a mock function with given fields: ctx, params
func (m *AccountRepository) GetByPhone(ctx context.Context, params models.AccountRequestByPhone) (*models.Account, error) {
	ret := m.Called(ctx, params)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(context.Context, models.AccountRequestByPhone) *models.Account); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.AccountRequestByPhone) error); ok {
		r1 = rf(ctx, params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (m *AccountRepository) List(ctx context.Context) ([]models.Account, error) {
	ret := m.Called(ctx)

	var r0 []models.Account
	if rf, ok := ret.Get(0).(func(context.Context) []models.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (m *AccountRepository) Update(ctx context.Context, params models.Account) error {
	ret := m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Account) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, params
func (m *AccountRepository) Delete(ctx context.Context, params models.AccountRequestById) error {
	ret := m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.AccountRequestById) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
