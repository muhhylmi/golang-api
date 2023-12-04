package usecases

import (
	"context"
	"golang-api/modules/cart/models/web"
	"golang-api/modules/cart/repositories"
	"golang-api/utils"

	"github.com/sirupsen/logrus"
)

const contextName = "modules.carts.usecase"

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.Repository
}
type Usecases interface {
	CreateCart(ctx context.Context, payload *web.RequestCreateCart) utils.Result
	GetAllCart(ctx context.Context, payload *web.RequestListCart) utils.Result
}

func NewUsecaseImpl(logger *logrus.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		logger:     logger,
		repository: repository,
	}
}
