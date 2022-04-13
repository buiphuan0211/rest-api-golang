package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env ENV

// InitDotEnv init params in .env file
func InitDotEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := GetEnvString("APP_PORT")
	database := Database{URI: GetEnvString("DB_URI"), Name: GetEnvString("DB_Name")}

	env = ENV{
		AppPort:  appPort,
		Database: database,
	}
}

// GetEnvString ...
func GetEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv return .env data
func GetEnv() *ENV {
	return &env
}
