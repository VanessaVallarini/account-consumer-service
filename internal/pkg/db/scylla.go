package db

import (
	"account-consumer-service/internal/models"
	"context"

	"github.com/gocql/gocql"
)

type Scylla struct {
	Session *gocql.Session
}

func NewScylla(c *models.DatabaseConfig) *Scylla {
	cluster := gocql.NewCluster(c.DatabaseHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.DatabaseUser,
		Password: c.DatabasePassword,
	}
	cluster.Keyspace = c.DatabaseKeyspace
	cluster.ConnectTimeout = cluster.ConnectTimeout * 5

	session, err := cluster.CreateSession()
	if err != nil {
		return nil
	}

	return &Scylla{
		Session: session,
	}
}

func (i *Scylla) Insert(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (i *Scylla) GetById(stmt string, ctx context.Context, values ...interface{}) *gocql.Query {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Consistency(gocql.One)
}

func (i *Scylla) List(stmt string, ctx context.Context) *gocql.Iter {
	q := i.Session.Query(stmt).WithContext(ctx)
	return q.Iter()
}

func (i *Scylla) Update(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (i *Scylla) Delete(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}
