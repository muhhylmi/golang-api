package repositories

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/modules/cart/models/web"
	"golang-api/utils/database"
	"golang-api/utils/logger"
)

type RepositoryImpl struct {
	Logger *logger.Logger
	DB     *database.DBService
}

type Repository interface {
	Save(cart *domain.Cart) (*domain.Cart, error)
	FindAll(payload *web.RequestListCart) ([]*domain.Cart, error)
}

func NewRepositoryImpl(logger *logger.Logger, db *database.DBService) Repository {
	return &RepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
