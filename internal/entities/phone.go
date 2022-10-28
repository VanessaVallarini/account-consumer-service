package entities

import "github.com/gocql/gocql"

type Phone struct {
	Id          gocql.UUID
	CountryCode string
	AreaCode    string
	Number      string
}
