package repositories

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/modules/cart/models/web"

	"gorm.io/gorm/clause"
)

func (repository *RepositoryImpl) Save(cart *domain.Cart) (*domain.Cart, error) {
	result := repository.DB.Gorm.Create(&cart)
	return cart, result.Error
}

func (repository *RepositoryImpl) FindAll(payload *web.RequestListCart) ([]*domain.Cart, error) {
	var carts []*domain.Cart
	tx := repository.DB.Gorm.
		Where("status = ?", payload.Status).
		Preload(clause.Associations).
		Preload("Details.Books").
		Find(&carts)
	return carts, tx.Error
}
