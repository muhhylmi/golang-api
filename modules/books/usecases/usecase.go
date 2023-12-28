package usecases

import (
	"context"
	"golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	"golang-api/utils"

	"github.com/sirupsen/logrus"
)

const contextName = "modules.books.usecase"

type UsecaseImpl struct {
	logger     *logrus.Logger
	repository repositories.Repository
}
type Usecases interface {
	CreateBook(ctx context.Context, payload *web.RequestCreateBook) utils.Result
	GetBook(ctx context.Context) utils.Result
	UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) utils.Result
	DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) utils.Result
	GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) utils.Result
	GetBookSheetData(ctx context.Context) utils.Result
}

func NewUsecaseImpl(logger *logrus.Logger, repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		logger:     logger,
		repository: repository,
	}
}
