package database

import (
	"golang-api/utils/logger"
)

func NewDatabase(params *DBServiceVar) (*DBService, error) {
	if params.Logger == nil {
		params.Logger = logger.Newlogger()
	}
	l := params.Logger.LogWithContext("dbConnection", "NewDatabase")

	pgDb, err := InitPostgres(params)
	if err != nil {
		l.Error("failed to init postgres")
		return nil, err
	}

	return &DBService{
		Gorm: pgDb,
	}, nil
}
