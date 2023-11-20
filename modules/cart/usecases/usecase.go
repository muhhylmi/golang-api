package usecases

import (
	"context"
	"golang-api/modules/cart/models/web"
	"golang-api/utils"
)

const contextName = "modules.carts.usecase"

type Usecases interface {
	CreateCart(ctx context.Context, payload *web.RequestCreateCart) utils.Result
	GetAllCart(ctx context.Context, payload *web.RequestListCart) utils.Result
}
