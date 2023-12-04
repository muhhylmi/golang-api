package repositories

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/modules/cart/models/web"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	logger *logrus.Logger
	db     *gorm.DB
}

type Repository interface {
	Save(cart *domain.Cart) (*domain.Cart, error)
	FindAll(payload *web.RequestListCart) ([]*domain.Cart, error)
}

func NewRepositoryImpl(logger *logrus.Logger, db *gorm.DB) Repository {
	return &RepositoryImpl{
		logger: logger,
		db:     db,
	}
}
