package scylla

import (
	"account-consumer-service/internal/entities"
	"context"
	"fmt"

	"github.com/gocql/gocql"
)

func NewScylla(c *entities.DatabaseConfig) *IScylla {
	cluster := gocql.NewCluster(c.DatabaseHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.DatabaseUser,
		Password: c.DatabasePassword,
	}
	cluster.Keyspace = c.DatabaseKeyspace
	cluster.ConnectTimeout = cluster.ConnectTimeout * 5
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	return &IScylla{
		Session: session,
	}
}

type IScylla struct {
	Session *gocql.Session
}

func (i *IScylla) Insert(stmt string, ctx context.Context, values ...interface{}) error {
	q := i.Session.Query(stmt, values).WithContext(ctx)
	return q.Exec()
}
