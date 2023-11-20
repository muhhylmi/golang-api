package repositories

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/modules/cart/models/web"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repository *RepositoryImpl) FindAll(payload *web.RequestListCart) ([]*domain.Cart, error) {
	var carts []*domain.Cart
	tx := repository.db.
		Where("status = ?", payload.Status).
		Preload(clause.Associations).
		Preload("Details.Books").
		Find(&carts)
	return carts, tx.Error
}
