package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Keyspace string
}

func NewConfig() *DatabaseConfig {
	viperConfig := initConfig()

	return &DatabaseConfig{
		Host:     viperConfig.GetString("DATEBASE_HOST"),
		Username: viperConfig.GetString("DATABASE_USERNAME"),
		Password: viperConfig.GetString("DATEBASE_PASSWORD"),
		Keyspace: viperConfig.GetString("DATEBASE_KEYSPACE"),
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
