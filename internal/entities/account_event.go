package entities

type EventAction string

const (
	AccountSubject = "account-consumer-service.pipeline.AccountEvent"
)

type AccountEvent struct {
	Id string `json:"id" avro:"id"`

	UserId string `json:"user_id" avro:"user_id"`
	Name   string `json:"name" avro:"name"`
	Email  string `json:"email" avro:"email"`

	AddressId   string `json:"address_id" avro:"address_id"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	District    string `json:"district" avro:"district"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`

	PhoneId     string `json:"phone_id" avro:"phone_id"`
	CountryCode string `json:"country_code" avro:"country_code"`
	AreaCode    string `json:"area_code" avro:"area_code"`
	Number      string `json:"number" avro:"number"`
	Op          string `json:"op" avro:"op"`
}
