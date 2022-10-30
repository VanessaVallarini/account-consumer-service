package entities

import "github.com/gocql/gocql"

type User struct {
	Id        gocql.UUID `json:"id"`
	AddressId string     `json:"address_id"`
	PhoneId   string     `json:"phone_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
}
