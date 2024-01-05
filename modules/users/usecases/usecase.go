package usecases

import (
	"context"
	"golang-api/modules/users/models/web"
	"golang-api/modules/users/repositories"
	"golang-api/utils/config"
	"golang-api/utils/logger"
	"golang-api/utils/wrapper"
)

const contextName = "modules.users.usecase"

type UsecaseImpl struct {
	Logger     *logger.Logger
	Repository repositories.Repository
	Config     *config.Configurations
}

type Usecases interface {
	CreateUser(ctx context.Context, payload *web.RequestCreateUser) wrapper.Result
	LoginUser(ctx context.Context, payload *web.RequestLogin) wrapper.Result
}

func NewUsecaseImpl(config *config.Configurations, logger *logger.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	}
}
