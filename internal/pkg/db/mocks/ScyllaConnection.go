package mocks

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/maraino/go-mock"
)

type ScyllaInterface interface {
	Insert(ctx context.Context, stmt string, values ...interface{}) error
	GetById(ctx context.Context, stmt string, values ...interface{}) *gocql.Query
	List(ctx context.Context, stmt string) *gocql.Iter
	Update(ctx context.Context, stmt string, values ...interface{}) error
	Delete(ctx context.Context, stmt string, values ...interface{}) error
}

type Scylla struct {
	mock.Mock
}

func NewScylla() *Scylla {
	return &Scylla{}
}

func (m *Scylla) Insert(ctx context.Context, stmt string, values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, stmt)
	_ca = append(_ca, values...)
	ret := m.Called(_ca...)

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, stmt, values...)
	} else {
		r1 = ret.Error(1)
	}

	return r1
}

func (m *Scylla) GetById(ctx context.Context, stmt string, values ...interface{}) *gocql.Query {
	return nil
}

func (m *Scylla) List(ctx context.Context, stmt string) *gocql.Iter {

	return nil
}

func (m *Scylla) Update(ctx context.Context, stmt string, values ...interface{}) error {
	q := m.Called(stmt, values)
	return q.Error(0)
}

func (m *Scylla) Delete(ctx context.Context, stmt string, values ...interface{}) error {
	q := m.Called(stmt, values)
	return q.Error(0)
}
