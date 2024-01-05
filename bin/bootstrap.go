package main

import (
	"golang-api/utils"
	"golang-api/utils/app"
	"golang-api/utils/config"
	"golang-api/utils/database"
	"golang-api/utils/logger"

	"github.com/labstack/echo/v4"
)

func Init() *app.App {
	config := config.GetConfig()
	logger := logger.Newlogger()
	e := echo.New()
	validator := utils.NewValidationUtil()

	PostgresUsername := config.DB_POSTGRES_USERNAME
	PostgresPassword := config.DB_POSTGRES_PASSWORD
	PostgresHost := config.DB_POSTGRES_HOST
	PostgresPort := config.DB_POSTGRES_PORT
	PostgresDBName := config.DB_POSTGRES_NAME
	PostgresSchema := config.DB_POSTGRES_SCHEMA
	dsn := "host=" + PostgresHost + " user=" + PostgresUsername + " password=" + PostgresPassword + " dbname=" + PostgresDBName + " port=" + PostgresPort + " sslmode=disable TimeZone=Asia/Shanghai" + " search_path=" + PostgresSchema

	db, err := database.NewDatabase(&database.DBServiceVar{
		Logger:      logger,
		PostgresUri: &dsn,
	})
	if err != nil {
		panic(err)
	}

	return &app.App{
		DBService:    db,
		Apps:         e,
		Validator:    validator,
		Logger:       logger,
		GlobalConfig: config,
	}

}
