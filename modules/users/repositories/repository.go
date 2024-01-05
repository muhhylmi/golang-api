package repositories

import (
	"golang-api/modules/users/models/domain"
	"golang-api/utils/database"
	"golang-api/utils/logger"
)

const contextName = "modules.users.repository"

type RepositoryImpl struct {
	Logger *logger.Logger
	DB     *database.DBService
}
type Repository interface {
	Save(user *domain.Users) (*domain.Users, error)
	FindByUsername(username string) (*domain.Users, error)
	FindById(Id string) (*domain.Users, error)
}

func NewRepositoryImpl(logger *logger.Logger, db *database.DBService) Repository {
	return &RepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
