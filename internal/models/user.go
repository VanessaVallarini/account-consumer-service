package models

type UserDBModel struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequestById struct {
	Id string `json:"id" validate:"required"`
}

type UserRequestByEmail struct {
	Email string `json:"email" validate:"required"`
}
