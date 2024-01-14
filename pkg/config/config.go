package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	DBhost string
	DBuser string
	DBpassword string
	DBport string
	DBname string
}

func Init() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	config := &Config{
		Port: viper.GetString("PORT"),
		DBhost: viper.GetString("DB_HOST"),
		DBuser: viper.GetString("DB_USER"),
		DBpassword: viper.GetString("DB_PASSWORD"),
		DBname: viper.GetString("DB_NAME"),
		DBport: viper.GetString("DB_PORT"),
	}

	return config;
}