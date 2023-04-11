package config

import (
	"account-consumer-service/internal/models"
	"account-consumer-service/internal/pkg/utils"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func NewConfig() (*models.Config, error) {
	viperConfig, err := initConfig()
	if err != nil {
		utils.Logger.Error("failed to read config file: %v", err)
		return nil, err
	}

	return &models.Config{
		AppName:          viperConfig.GetString("APP_NAME"),
		ServerHost:       viperConfig.GetString("SERVER_HOST"),
		HealthServerHost: viperConfig.GetString("HEALTH_SERVER_HOST"),
		Database:         buildDatabaseConfig(viperConfig),
		Kafka:            buildKafkaClientConfig(viperConfig),
	}, nil
}

func initConfig() (*viper.Viper, error) {
	config := viper.New()

	config.SetConfigType("yml")
	config.SetConfigName("configuration")
	config.AddConfigPath("internal/config/")

	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config.AutomaticEnv()

	return config, nil
}

func buildDatabaseConfig(viperConfig *viper.Viper) *models.DatabaseConfig {
	return &models.DatabaseConfig{
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

func buildKafkaClientConfig(config *viper.Viper) *models.KafkaConfig {
	return &models.KafkaConfig{
		ClientId:               config.GetString("KAFKA_CLIENT_ID"),
		Hosts:                  cast.ToStringSlice(config.GetString("KAFKA_HOSTS")),
		UseAuthentication:      config.GetBool("KAFKA_USE_AUTHENTICATION"),
		SaslMechanism:          config.GetString("KAFKA_SASL_MECHANISM"),
		EnableTLS:              config.GetBool("KAFKA_ENABLE_TLS"),
		SchemaRegistryHost:     config.GetString("KAFKA_SCHEMA_REGISTRY_HOST"),
		User:                   config.GetString("KAFKA_USER"),
		Password:               config.GetString("KAFKA_PASSWORD"),
		SchemaRegistryUser:     config.GetString("KAFKA_SCHEMA_REGISTRY_USER"),
		SchemaRegistryPassword: config.GetString("KAFKA_SCHEMA_REGISTRY_PASSWORD"),
		DlqTopic:               cast.ToStringSlice(config.GetString("KAFKA_QLD_TOPIC_NAME")),
		ConsumerTopic:          cast.ToStringSlice(config.GetString("KAFKA_CONSUMER_TOPIC_NAME")),
		ConsumerGroup:          config.GetString("KAFKA_CONSUMER_GROUP"),
	}
}
