package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DatabaseURL string
}

func LoadPath() (*DBConfig, error) {
	viper.SetDefault("DATABASE_URL", "clothing_pair_project")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	config := &DBConfig{
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}

	if config.DatabaseURL == "" {
		return nil, fmt.Errorf("database configuration fields are required")
	}

	return config, nil
}
