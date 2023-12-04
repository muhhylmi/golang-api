package repositories

import (
	"golang-api/modules/users/models/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const contextName = "modules.users.repository"

type RepositoryImpl struct {
	logger *logrus.Logger
	db     *gorm.DB
}
type Repository interface {
	Save(user *domain.Users) (*domain.Users, error)
	FindByUsername(username string) (*domain.Users, error)
	FindById(Id string) (*domain.Users, error)
}

func NewRepositoryImpl(logger *logrus.Logger, db *gorm.DB) Repository {
	return &RepositoryImpl{
		logger: logger,
		db:     db,
	}
}
