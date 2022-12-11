package models

type Config struct {
	AppName          string
	ServerHost       string
	HealthServerHost string
	Database         *DatabaseConfig
	Kafka            *KafkaConfig
}

type DatabaseConfig struct {
	DatabaseUser                string
	DatabasePassword            string
	DatabaseKeyspace            string
	DatabaseHost                string
	DatabasePort                int
	DatabaseConnectionRetryTime int
	DatabaseRetryMinArg         int
	DatabaseRetryMaxArg         int
	DatabaseNumRetries          int
	DatabaseClusterTimeout      int
}

type KafkaConfig struct {
	ClientId               string
	Hosts                  []string
	SchemaRegistryHost     string
	Acks                   string
	Timeout                int
	UseAuthentication      bool
	EnableTLS              bool
	SaslMechanism          string
	User                   string
	Password               string
	SchemaRegistryUser     string
	SchemaRegistryPassword string
	EnableEvents           bool
	MaxMessageBytes        int
	RetryMax               int
	DlqTopic               string
	ConsumerTopic          []string
	ConsumerGroup          string
}
