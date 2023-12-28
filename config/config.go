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
	BASIC_AUTH_USERNAME  string
	BASIC_AUTH_PASSWORD  string
	JWT_SECRET_KEY       string

	SPREAD_SHEET_ID           string
	SHEET_KEY_TYPE            string
	SHEET_KEY_PROJECT_ID      string
	SHEET_KEY_PRIVATE_KEY_ID  string
	SHEET_KEY_PRIVATE_KEY     string
	SHEET_KEY_CLIENT_ID       string
	SHEET_KEY_CLIENT_EMAIL    string
	SHEET_KEY_AUTH_URI        string
	SHEET_KEY_TOKEN_URI       string
	SHEET_KEY_AUTH_PROVIDER   string
	SHEET_KEY_CLIENT_CERT_URI string
	SHEET_KEY_UNIV_DOMAIN     string
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
		BASIC_AUTH_USERNAME:  os.Getenv("BASIC_AUTH_USERNAME"),
		BASIC_AUTH_PASSWORD:  os.Getenv("BASIC_AUTH_PASSWORD"),
		JWT_SECRET_KEY:       os.Getenv("JWT_SECRET_KEY"),

		SPREAD_SHEET_ID:           os.Getenv("SPREAD_SHEET_ID"),
		SHEET_KEY_TYPE:            os.Getenv("SHEET_KEY_TYPE"),
		SHEET_KEY_PROJECT_ID:      os.Getenv("SHEET_KEY_PROJECT_ID"),
		SHEET_KEY_PRIVATE_KEY_ID:  os.Getenv("SHEET_KEY_PRIVATE_KEY_ID"),
		SHEET_KEY_PRIVATE_KEY:     os.Getenv("SHEET_KEY_PRIVATE_KEY"),
		SHEET_KEY_CLIENT_ID:       os.Getenv("SHEET_KEY_CLIENT_ID"),
		SHEET_KEY_CLIENT_EMAIL:    os.Getenv("SHEET_KEY_CLIENT_EMAIL"),
		SHEET_KEY_AUTH_URI:        os.Getenv("SHEET_KEY_AUTH_URI"),
		SHEET_KEY_TOKEN_URI:       os.Getenv("SHEET_KEY_TOKEN_URI"),
		SHEET_KEY_AUTH_PROVIDER:   os.Getenv("SHEET_KEY_AUTH_PROVIDER"),
		SHEET_KEY_CLIENT_CERT_URI: os.Getenv("SHEET_KEY_CLIENT_CERT_URI"),
		SHEET_KEY_UNIV_DOMAIN:     os.Getenv("SHEET_KEY_UNIV_DOMAIN"),
	}
}
