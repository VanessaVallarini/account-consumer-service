package db

import (
	"account-consumer-service/internal/models"
	"context"

	"github.com/gocql/gocql"
)

type ScyllaInterface interface {
	Insert(ctx context.Context, stmt string, values ...interface{}) error
	ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error
	List(ctx context.Context, stmt string) *gocql.Iter
	Update(ctx context.Context, stmt string, values ...interface{}) error
	Delete(ctx context.Context, stmt string, values ...interface{}) error
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

func (s *Scylla) Insert(ctx context.Context, stmt string, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error {
	q := s.Session.Query(stmt, arguments...).WithContext(ctx)
	return q.MapScan(results)
}

func (s *Scylla) List(ctx context.Context, stmt string) *gocql.Iter {
	q := s.Session.Query(stmt).WithContext(ctx)
	return q.Iter()
}

func (s *Scylla) Update(ctx context.Context, stmt string, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) Delete(ctx context.Context, stmt string, values ...interface{}) error {
	q := s.Session.Query(stmt, values...).WithContext(ctx)
	return q.Exec()
}
