package usecases

import (
	"context"
	"golang-api/modules/users/models/web"
	"golang-api/utils"
)

const contextName = "modules.users.usecase"

type Usecases interface {
	CreateUser(ctx context.Context, payload *web.RequestCreateUser) utils.Result
	LoginUser(ctx context.Context, payload *web.RequestLogin) utils.Result
}
