package repositories

import "golang-api/modules/cart/models/domain"

type Repository interface {
	Save(cart *domain.Cart) (*domain.Cart, error)
}
