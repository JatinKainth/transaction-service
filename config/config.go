package config

import (
	"fmt"
	"transaction_service/pkg/db"

	"github.com/spf13/viper"
)

type Config struct {
	Database db.DatabaseConfig `json:"database" toml:"database"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("credentials")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
