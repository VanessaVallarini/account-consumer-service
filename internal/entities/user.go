package entities

import "github.com/gocql/gocql"

type User struct {
	Id        gocql.UUID
	AddressId string
	PhoneId   string
	Name      string
	Email     string
}
