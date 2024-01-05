package usecases

import (
	"context"
	"golang-api/modules/cart/models/web"
	"golang-api/utils/wrapper"
)

func (usecase *UsecaseImpl) CreateCart(ctx context.Context, payload *web.RequestCreateCart) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "CreateCart")
	var result wrapper.Result
	cartData := payload.ToModel()
	book, err := usecase.Repository.Save(cartData)
	if err != nil {
		log.Error("Cannot Create Cart")
		error := wrapper.NewBadRequest("Cannot Create Cart")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) GetAllCart(ctx context.Context, payload *web.RequestListCart) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "GetAllCart")
	var result wrapper.Result
	carts, err := usecase.Repository.FindAll(payload)
	if err != nil {
		log.Error("Book is not found")
		error := wrapper.NewBadRequest("Book Is Not Found")
		result.Error = error
		return result
	}
	responseData := web.ToResponseCart(carts)
	result.Data = responseData
	return result
}
