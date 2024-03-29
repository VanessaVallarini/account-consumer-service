package db

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"
	"context"

	"github.com/gocql/gocql"
)

type IScylla interface {
	Insert(ctx context.Context, stmt string, arguments ...interface{}) error
	ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error
	Delete(ctx context.Context, stmt string, arguments ...interface{}) error
	Close()
}

type Scylla struct {
	session *gocql.Session
}

func NewScylla(c *models.DatabaseConfig) (*Scylla, error) {
	cluster := gocql.NewCluster(c.DatabaseHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.DatabaseUser,
		Password: c.DatabasePassword,
	}
	cluster.Keyspace = c.DatabaseKeyspace
	cluster.ConnectTimeout = cluster.ConnectTimeout * 5
	cluster.ProtoVersion = 4

	session, err := cluster.CreateSession()
	if err != nil {
		utils.Logger.Error("failed to create session: %v", err)
		return nil, err
	}

	return &Scylla{
		session: session,
	}, nil
}

func (s *Scylla) Insert(ctx context.Context, stmt string, arguments ...interface{}) error {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.MapScan(results)
}

func (s *Scylla) Delete(ctx context.Context, stmt string, arguments ...interface{}) error {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) Close() {
	s.session.Close()
}
