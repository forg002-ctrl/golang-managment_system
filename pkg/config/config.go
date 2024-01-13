package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func Init() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	config := &Config{
		Port: viper.GetString("PORT"),
	}

	return config;
}