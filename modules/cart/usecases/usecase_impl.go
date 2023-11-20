package usecases

import (
	"context"
	"golang-api/modules/cart/models/web"
	"golang-api/modules/cart/repositories"
	"golang-api/utils"

	"github.com/sirupsen/logrus"
)

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.Repository
}

func NewUsecaseImpl(logger *logrus.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		logger:     logger,
		repository: repository,
	}
}

func (usecase *UsecaseImpl) CreateCart(ctx context.Context, payload *web.RequestCreateCart) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateCart")
	var result utils.Result
	cartData := payload.ToModel()
	book, err := usecase.repository.Save(cartData)
	if err != nil {
		log.Error("Cannot Create Cart")
		error := utils.NewBadRequest("Cannot Create Cart")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) GetAllCart(ctx context.Context, payload *web.RequestListCart) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "GetAllCart")
	var result utils.Result
	carts, err := usecase.repository.FindAll(payload)
	if err != nil {
		log.Error("Book is not found")
		error := utils.NewBadRequest("Book Is Not Found")
		result.Error = error
		return result
	}
	responseData := web.ToResponseCart(carts)
	result.Data = responseData
	return result
}
