package db

import (
	"account-consumer-service/internal/models"
	"context"

	"github.com/gocql/gocql"
)

type ScyllaConnection interface {
	Insert(stmt string, ctx context.Context, values ...interface{}) error
	GetById(stmt string, ctx context.Context, values ...interface{}) *gocql.Query
	List(stmt string, ctx context.Context) *gocql.Iter
	Update(stmt string, ctx context.Context, values ...interface{}) error
	Delete(stmt string, ctx context.Context, values ...interface{}) error
}

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

func (s *Scylla) Insert(stmt string, ctx context.Context, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) GetById(stmt string, ctx context.Context, values ...interface{}) *gocql.Query {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Consistency(gocql.One)
}

func (s *Scylla) List(stmt string, ctx context.Context) *gocql.Iter {
	q := s.Session.Query(stmt).WithContext(ctx)
	return q.Iter()
}

func (s *Scylla) Update(stmt string, ctx context.Context, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) Delete(stmt string, ctx context.Context, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}
