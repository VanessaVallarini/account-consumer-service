package config

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config struct {
	AppName          string
	ServerHost       string
	HealthServerHost string
	DatabaseConnStr  *DatabaseConfig
	Kafka            *KafkaConfig
}

// DatabaseConfigs holds all the database connection parameters
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

// KafkaConfigs holds all the kafka connection parameters
type KafkaConfig struct {
	ClientId                             string
	Hosts                                []string
	SchemaRegistryHost                   string
	Acks                                 string
	Timeout                              int
	UseAuthentication                    bool
	EnableTLS                            bool
	SaslMechanism                        string
	User                                 string
	Password                             string
	SchemaRegistryUser                   string
	SchemaRegistryPassword               string
	EnableEvents                         bool
	MaxMessageBytes                      int
	RetryMax                             int
	DlqTopic                             string
	ConsumerTopic                        string
	ConsumerTopicStrategiesManagement    string
	ConsumerTopicStrategiesManagementDLQ string
	ConsumerGroup                        string
}

func NewConfig() *Config {
	viperConfig := initConfig()

	return &Config{
		AppName:          viperConfig.GetString("APP_NAME"),
		ServerHost:       viperConfig.GetString("SERVER_HOST"),
		HealthServerHost: viperConfig.GetString("HEALTH_SERVER_HOST"),
		DatabaseConnStr:  buildDatabaseConfig(viperConfig),
		Kafka:            buildKafkaClientConfig(viperConfig),
	}
}

func initConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigType("yml")
	config.SetConfigName("configuration")
	config.AddConfigPath("internal/config/")

	err := config.ReadInConfig()
	if err != nil {
		fmt.Println(err, "failed to read config file")
	}

	config.AutomaticEnv()

	return config
}

func buildDatabaseConfig(viperConfig *viper.Viper) *DatabaseConfig {
	return &DatabaseConfig{
		DatabaseUser:                viperConfig.GetString("DATABASE_USER"),
		DatabasePassword:            viperConfig.GetString("DATEBASE_PASSWORD"),
		DatabaseKeyspace:            viperConfig.GetString("DATEBASE_KEYSPACE"),
		DatabaseHost:                viperConfig.GetString("DATEBASE_HOST"),
		DatabasePort:                viperConfig.GetInt("DATEBASE_PORT"),
		DatabaseConnectionRetryTime: viperConfig.GetInt("DATEBASE_CONNECTION_RETRY_TIME"),
		DatabaseRetryMinArg:         viperConfig.GetInt("DATEBASE_RETRY_MIN"),
		DatabaseRetryMaxArg:         viperConfig.GetInt("DATEBASE_RETRY_MAX"),
		DatabaseNumRetries:          viperConfig.GetInt("DATEBASE_NUM_RETRIES"),
		DatabaseClusterTimeout:      viperConfig.GetInt("DATEBASE_CLUSTER_TIMEOUT"),
	}
}

func buildKafkaClientConfig(config *viper.Viper) *KafkaConfig {
	return &KafkaConfig{
		ClientId:                             config.GetString("KAFKA_CLIENT_ID"),
		Hosts:                                cast.ToStringSlice(config.GetString("KAFKA_HOSTS")),
		SchemaRegistryHost:                   config.GetString("KAFKA_SCHEMA_REGISTRY_HOST"),
		Acks:                                 config.GetString("KAFKA_ACKS"),
		Timeout:                              config.GetInt("KAFKA_TIMEOUT"),
		UseAuthentication:                    config.GetBool("KAFKA_HAS_AUTH"),
		EnableTLS:                            config.GetBool("KAFKA_ENABLE_TLS"),
		SaslMechanism:                        config.GetString("KAFKA_SASL_MECHANISM"),
		User:                                 config.GetString("KAFKA_USER"),
		Password:                             config.GetString("KAFKA_PASSWORD"),
		SchemaRegistryUser:                   config.GetString("KAFKA_SCHEMA_REGISTRY_USER"),
		SchemaRegistryPassword:               config.GetString("KAFKA_SCHEMA_REGISTRY_PASSWORD"),
		EnableEvents:                         config.GetBool("KAFKA_ENABLE_EVENTS"),
		MaxMessageBytes:                      config.GetInt("KAFKA_MAX_MESSAGE_BYTES"),
		RetryMax:                             config.GetInt("KAFKA_RETRY_MAX"),
		DlqTopic:                             config.GetString("KAFKA_DLQ_TOPIC"),
		ConsumerTopic:                        config.GetString("KAFKA_CONSUMER_TOPIC"),
		ConsumerTopicStrategiesManagement:    config.GetString("KAFKA_CONSUMER_TOPIC_STRATEGIES_MANAGEMENT"),
		ConsumerTopicStrategiesManagementDLQ: config.GetString("KAFKA_CONSUMER_TOPIC_STRATEGIES_MANAGEMENT_DLQ"),
		ConsumerGroup:                        config.GetString("KAFKA_CONSUMER_GROUP"),
	}
}
