package usecases

import (
	"context"
	"golang-api/modules/cart/models/web"
	"golang-api/modules/cart/repositories"
	"golang-api/utils/config"
	"golang-api/utils/logger"
	"golang-api/utils/wrapper"
)

const contextName = "modules.carts.usecase"

type UsecaseImpl struct {
	Logger     *logger.Logger
	Repository repositories.Repository
	Config     *config.Configurations
}
type Usecases interface {
	CreateCart(ctx context.Context, payload *web.RequestCreateCart) wrapper.Result
	GetAllCart(ctx context.Context, payload *web.RequestListCart) wrapper.Result
}

func NewUsecaseImpl(config *config.Configurations, logger *logger.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	}
}
