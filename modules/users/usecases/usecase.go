package usecases

import (
	"context"
	"golang-api/modules/users/models/web"
	"golang-api/modules/users/repositories"
	"golang-api/utils"

	"github.com/sirupsen/logrus"
)

const contextName = "modules.users.usecase"

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.Repository
}

type Usecases interface {
	CreateUser(ctx context.Context, payload *web.RequestCreateUser) utils.Result
	LoginUser(ctx context.Context, payload *web.RequestLogin) utils.Result
}

func NewUsecaseImpl(logger *logrus.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		logger:     logger,
		repository: repository,
	}
}
