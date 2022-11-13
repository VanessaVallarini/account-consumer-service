package models

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
	Command     string `json:"command"`
}

type AccountRequestById struct {
	Id string `json:"id" validate:"required"`
}

type AccountRequestByEmail struct {
	Email string `json:"email" validate:"required"`
}

type AccountRequestByPhone struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
}
