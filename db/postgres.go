package db

import (
	"golang-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PostgresDB *gorm.DB

func InitPostgres() *gorm.DB {
	config := config.GetConfig()

	PostgresUsername := config.DB_POSTGRES_USERNAME
	PostgresPassword := config.DB_POSTGRES_PASSWORD
	PostgresHost := config.DB_POSTGRES_HOST
	PostgresPort := config.DB_POSTGRES_PORT
	PostgresDBName := config.DB_POSTGRES_NAME

	dsn := "host=" + PostgresHost + " user=" + PostgresUsername + " password=" + PostgresPassword + " dbname=" + PostgresDBName + " port=" + PostgresPort + " sslmode=disable TimeZone=Asia/Shanghai"
	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public",
		},
	})
	if err != nil {
		panic(err)
	}
	return PostgresDB
}
