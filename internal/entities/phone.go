package entities

import "github.com/gocql/gocql"

type Phone struct {
	Id          gocql.UUID `json:"id"`
	CountryCode string     `json:"country_code"`
	AreaCode    string     `json:"area_code"`
	Number      string     `json:"number"`
}
