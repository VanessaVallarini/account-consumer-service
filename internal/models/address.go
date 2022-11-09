package models

import "github.com/gocql/gocql"

type Address struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
}

type AddressDbo struct {
	Id           gocql.UUID
	Alias        string
	City         string
	District     string
	Public_place string
	Zip_code     string
}

type ListAddressDbo struct {
	Id           string
	Alias        string
	City         string
	District     string
	Public_place string
	Zip_code     string
}

type AddressRequestById struct {
	Id string `json:"id" validate:"required"`
}
