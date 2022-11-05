package scylla

import (
	"context"

	"github.com/gocql/gocql"
)

func (i *IScylla) Insert(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (i *IScylla) GetById(stmt string, ctx context.Context, values ...interface{}) *gocql.Query {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Consistency(gocql.One)
}

func (i *IScylla) List(stmt string, ctx context.Context) *gocql.Iter {
	q := i.Session.Query(stmt).WithContext(ctx)
	return q.Iter()
}

func (i *IScylla) Update(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (i *IScylla) Delete(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}
