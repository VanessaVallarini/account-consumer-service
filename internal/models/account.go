package models

type AccountCreate struct {
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

type AccountUpdate struct {
	Email         string `json:"email"`
	FullNumber    string `json:"full_number"`
	Alias         string `json:"alias"`
	City          string `json:"city"`
	DateTime      string `json:"date_time"`
	District      string `json:"district"`
	Name          string `json:"name"`
	PublicPlace   string `json:"public_place"`
	Status        string `json:"status"`
	ZipCode       string `json:"zip_code"`
	OldEmail      string `json:"old_email"`
	OldFullNumber string `json:"old_full_number"`
}

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

type AccountRequestBy struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	FullNumber string `json:"full_number"`
}

type AccountRequestByEmailAndFullNumber struct {
	Email      string `json:"email"`
	FullNumber string `json:"full_number"`
}

type AccountRequestByEmail struct {
	Email string `json:"email"`
}

type AccountRequestByFullNumber struct {
	FullNumber string `json:"full_number"`
}

type AccountRequestById struct {
	Id string `json:"id"`
}
