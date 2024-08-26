package main

import (
	"github.com/spf13/viper"
	"log"
)

// Config structure to hold the environment variables.
type Config struct {
	Host             string
	Port             string
	Environment      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

// LoadConfig loads environment variables using Viper and returns a Config structure.
func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Reads environment variables

	// Default environment variables
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("ENVIRONMENT", "development")

	// Load the .env file if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("The .env configuration file was not found or there was an error reading it: %v", err)
	}

	config := &Config{
		Host:             viper.GetString("HOST"),
		Port:             viper.GetString("PORT"),
		Environment:      viper.GetString("ENVIRONMENT"),
		DatabaseHost:     viper.GetString("DATABASE_HOST"),
		DatabasePort:     viper.GetString("DATABASE_PORT"),
		DatabaseName:     viper.GetString("DATABASE_NAME"),
		DatabaseUser:     viper.GetString("DATABASE_USER"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
	}

	return config
}
