package repositories

import (
	"golang-api/modules/cart/models/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewRepositoryImpl(logger *logrus.Logger, db *gorm.DB) Repository {
	return &RepositoryImpl{
		logger: logger,
		db:     db,
	}
}

func (repository *RepositoryImpl) Save(cart *domain.Cart) (*domain.Cart, error) {
	result := repository.db.Create(&cart)
	return cart, result.Error
}
