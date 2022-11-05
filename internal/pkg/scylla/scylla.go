package scylla

import (
	"account-consumer-service/internal/models"

	"github.com/gocql/gocql"
)

type IScylla struct {
	Session *gocql.Session
}

func NewScylla(c *models.DatabaseConfig) *IScylla {
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

	return &IScylla{
		Session: session,
	}
}
