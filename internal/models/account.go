package models

type Account struct {
	Email       string `json:"email"`
	FullNumber  string `json:"full_number"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	DateTime    string `json:"date_time"`
	District    string `json:"district"`
	Name        string `json:"name"`
	PublicPlace string `json:"public_place"`
	Status      string `json:"status"`
	ZipCode     string `json:"zip_code"`
}

type AccountRequestByEmail struct {
	Email string `json:"email"`
}
