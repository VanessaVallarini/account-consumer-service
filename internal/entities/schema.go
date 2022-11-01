package entities

const (
	UserSubject = "com.account.producer"
	UserAvro    = `{
		"type": "record",
		"name": "UserAccount",
		"namespace": "com.account.producer",
		"fields": [
			{ "name": "name", "type": "string" },
			{ "name": "email", "type": "string" },
		   ]
	   }`
)
