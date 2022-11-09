package models

const (
	AccountSubject = "com.account.producer"
	AccountAvro    = `{
		"type":"record",
		"name":"Account",
		"namespace":"com.account.producer",
		"fields":[
			{
				"name":"id",
				"type":"string"
			 },
			 {
				"name":"name",
				"type":"string"
			 },
			 {
				"name":"email",
				"type":"string"
			 },
			 {
				"name":"alias",
				"type":"string"
			 },
			 {
				"name":"city",
				"type":"string"
			 },
			 {
				"name":"district",
				"type":"string"
			 },
			 {
				"name":"public_place",
				"type":"string"
			 },
			 {
				"name":"zip_code",
				"type":"string"
			 },
			 {
				"name":"country_code",
				"type":"string"
			 },
			 {
				"name":"area_code",
				"type":"string"
			 },
			 {
				"name":"number",
				"type":"string"
			 },
			 {
				"name":"command",
				"type":"string"
			 }		   
		]
	 }`
)

type AccountEvent struct {
	Id          string `json:"id" avro:"id"`
	Name        string `json:"name" avro:"name"`
	Email       string `json:"email" avro:"email"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	District    string `json:"district" avro:"district"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
	CountryCode string `json:"country_code" avro:"country_code"`
	AreaCode    string `json:"area_code" avro:"area_code"`
	Number      string `json:"number" avro:"number"`
	Command     string `json:"command" avro:"command"`
}
