package db

import (
	"golang-api/config"
	"golang-api/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PostgresDB *gorm.DB

func InitPostgres(logger *logrus.Logger) *gorm.DB {
	log := utils.LogWithContext(logger, "dbConnection", "InitPostgres")
	config := config.GetConfig()

	PostgresUsername := config.DB_POSTGRES_USERNAME
	PostgresPassword := config.DB_POSTGRES_PASSWORD
	PostgresHost := config.DB_POSTGRES_HOST
	PostgresPort := config.DB_POSTGRES_PORT
	PostgresDBName := config.DB_POSTGRES_NAME
	PostgresSchema := config.DB_POSTGRES_SCHEMA

	dsn := "host=" + PostgresHost + " user=" + PostgresUsername + " password=" + PostgresPassword + " dbname=" + PostgresDBName + " port=" + PostgresPort + " sslmode=disable TimeZone=Asia/Shanghai" + " search_path=" + PostgresSchema
	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: PostgresSchema,
		},
	})
	if err != nil {
		log.Info("Connection Postgres is Failed")
		panic(err)
	}
	log.Info("Success connect to database")
	return PostgresDB
}
