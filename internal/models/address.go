package models

type AddressDBModel struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
	UserId      string `json:"user_id"`
}

type AddressRequestById struct {
	Id string `json:"id"`
}
