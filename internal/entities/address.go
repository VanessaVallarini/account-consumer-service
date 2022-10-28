package entities

import "github.com/gocql/gocql"

type Address struct {
	Id          gocql.UUID
	Alias       string
	City        string
	District    string
	PublicPlace string
	ZipCode     string
}
