package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	AppName string `mapstructure:"app_name"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	bindings := map[string]string{
		"database.password": "POSTGRES_PASSWORD",
		"database.host":     "POSTGRES_HOST",
		"database.port":     "POSTGRES_PORT",
		"database.user":     "POSTGRES_USERNAME",
		"database.name":     "POSTGRES_DBNAME",
	}

	for key, env := range bindings {
		if err := viper.BindEnv(key, env); err != nil {
			return nil, fmt.Errorf("failed to bind env %s to %s: %w", env, key, err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Config file not found, using env or defaults")
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	return &cfg, err
}
