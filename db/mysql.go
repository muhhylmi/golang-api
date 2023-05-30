package db

import (
	"database/sql"
	"golang-api/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	config := config.GetConfig()
	connectionString := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetConn() *sql.DB {
	return db
}
