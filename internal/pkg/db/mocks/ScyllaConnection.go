package mocks

import (
	"context"

	"github.com/maraino/go-mock"
)

type ScyllaInterface interface {
	Insert(ctx context.Context, stmt string, values ...interface{}) error
	ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error
	ScanMapSlice(ctx context.Context, stmt string, arguments ...interface{}) ([]map[string]interface{}, error)
	Update(ctx context.Context, stmt string, values ...interface{}) error
	Delete(ctx context.Context, stmt string, values ...interface{}) error
}

// Scylla is an autogenerated mock type for the RepositoryInterface type
type Scylla struct {
	mock.Mock
}

func NewScylla() *Scylla {
	return &Scylla{}
}

// Insert provides a mock function with given fields: ctx, stmt, params
func (m *Scylla) Insert(ctx context.Context, stmt string, values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, stmt)
	_ca = append(_ca, values...)
	ret := m.Called(_ca...).Error(0)

	return ret
}

// Scan implements Session.
func (m *Scylla) ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, stmt)
	_ca = append(_ca, results)
	_ca = append(_ca, arguments...)
	ret := m.Called(_ca...).Error(0)
	return ret
}

func (m *Scylla) ScanMapSlice(ctx context.Context, stmt string, arguments ...interface{}) ([]map[string]interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, stmt)
	_ca = append(_ca, arguments...)
	ret := m.Called(_ca...)

	var r0 []map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string) []map[string]interface{}); ok {
		r0 = rf(ctx, stmt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, stmt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (m *Scylla) Update(ctx context.Context, stmt string, values ...interface{}) error {
	q := m.Called(stmt, values)
	return q.Error(0)
}

func (m *Scylla) Delete(ctx context.Context, stmt string, values ...interface{}) error {
	q := m.Called(stmt, values)
	return q.Error(0)
}
