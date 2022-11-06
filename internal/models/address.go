package models

type Address struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
}

type AddressRequestById struct {
	Id string `json:"id" validate:"required"`
}
