package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB_USERNAME          string
	DB_PASSWORD          string
	DB_HOST              string
	DB_PORT              string
	DB_NAME              string
	DB_POSTGRES_USERNAME string
	DB_POSTGRES_PASSWORD string
	DB_POSTGRES_HOST     string
	DB_POSTGRES_PORT     string
	DB_POSTGRES_NAME     string
	DB_POSTGRES_SCHEMA   string
	HOST                 string
}

func GetConfig() *Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return &Configuration{
		DB_USERNAME:          os.Getenv("DB_USERNAME"),
		DB_PASSWORD:          os.Getenv("DB_PASSWORD"),
		DB_HOST:              os.Getenv("DB_HOST"),
		DB_PORT:              os.Getenv("DB_PORT"),
		DB_NAME:              os.Getenv("DB_NAME"),
		DB_POSTGRES_USERNAME: os.Getenv("DB_POSTGRES_USERNAME"),
		DB_POSTGRES_PASSWORD: os.Getenv("DB_POSTGRES_PASSWORD"),
		DB_POSTGRES_HOST:     os.Getenv("DB_POSTGRES_HOST"),
		DB_POSTGRES_PORT:     os.Getenv("DB_POSTGRES_PORT"),
		DB_POSTGRES_NAME:     os.Getenv("DB_POSTGRES_NAME"),
		DB_POSTGRES_SCHEMA:   os.Getenv("DB_POSTGRES_SCHEMA"),
		HOST:                 os.Getenv("HOST_PORT"),
	}
}
