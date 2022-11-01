package entities

import "github.com/gocql/gocql"

type Address struct {
	Id          gocql.UUID `json:"id"`
	Alias       string     `json:"alias"`
	City        string     `json:"city"`
	District    string     `json:"district"`
	PublicPlace string     `json:"public_place"`
	ZipCode     string     `json:"zip_code"`
}

type AddressRequestById struct {
	Id string `json:"id" validate:"required"`
}
