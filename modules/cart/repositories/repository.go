package repositories

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/modules/cart/models/web"
)

type Repository interface {
	Save(cart *domain.Cart) (*domain.Cart, error)
	FindAll(payload *web.RequestListCart) ([]*domain.Cart, error)
}
