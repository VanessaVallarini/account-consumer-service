package config

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.ifoodcorp.com.br/kafka-client-go/kafka"
)

type Config struct {
	AppName          string
	ServerHost       string
	HealthServerHost string
	DatabaseConnStr  *DatabaseConfig
	Kafka            kafka.ClientConfig
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

func buildKafkaClientConfig(config *viper.Viper) kafka.ClientConfig {
	return kafka.ClientConfig{
		UseAuthentication:      config.GetBool("KAFKA_HAS_AUTH"),
		EnableTLS:              true,
		Acks:                   "all",
		BalanceStrategy:        kafka.BalanceStrategyRange,
		Timeout:                config.GetInt("KAFKA_TIMEOUT"),
		ClientId:               config.GetString("KAFKA_CLIENT_ID"),
		SaslMechanism:          config.GetString("KAFKA_SASL_MECHANISM"),
		KafkaUser:              config.GetString("KAFKA_USER"),
		KafkaPassword:          config.GetString("KAFKA_PASSWORD"),
		KafkaAddresses:         cast.ToStringSlice(config.GetString("KAFKA_ADDRESS")),
		SchemaRegistryHost:     config.GetString("KAFKA_SCHEMA_REGISTRY_HOST"),
		SchemaRegistryUser:     config.GetString("KAFKA_SCHEMA_REGISTRY_USER"),
		SchemaRegistryPassword: config.GetString("KAFKA_SCHEMA_REGISTRY_PASSWORD"),
		EnableEvents:           config.GetBool("KAFKA_ENABLE_EVENTS"),
		ConsumerConfig: &kafka.ConsumerConfig{
			Group:             nil,
			MaxProcessingTime: 0,
		},
	}
}
