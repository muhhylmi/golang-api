package usecases

import (
	"context"
	"golang-api/modules/books/models/web"
	"golang-api/utils"
)

type Usecases interface {
	CreateBook(ctx context.Context, payload *web.RequestCreateBook) utils.Result
	GetBook(ctx context.Context) utils.Result
	UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) utils.Result
	DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) utils.Result
	GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) utils.Result
}
