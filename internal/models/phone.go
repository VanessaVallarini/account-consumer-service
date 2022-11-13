package models

type PhoneDBModel struct {
	Id          string `json:"id"`
	AreaCode    string `json:"area_code"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
	UserId      string `json:"user_id"`
}

type PhoneRequestById struct {
	Id string `json:"id" validate:"required"`
}
