package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password int
		DBName   string
	}

	Server struct {
		Port int
	}
}

var Cfg Config

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}
}
