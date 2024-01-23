package usecases

import (
	"context"
	"golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	"golang-api/utils/config"
	googlesheet "golang-api/utils/google_sheet"
	utils_grpc "golang-api/utils/grpc"
	"golang-api/utils/logger"
	"golang-api/utils/wrapper"
)

const contextName = "modules.books.usecase"

type UsecaseImpl struct {
	Logger       *logger.Logger
	Repository   repositories.Repository
	Config       *config.Configurations
	Gsheet       googlesheet.GoogleSheetServiceInterface
	GrpcServices utils_grpc.GrpcServicesInterface
}
type Usecases interface {
	CreateBook(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result
	GetBook(ctx context.Context) wrapper.Result
	UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) wrapper.Result
	DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) wrapper.Result
	GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) wrapper.Result
	GetBookSheetData(ctx context.Context, payload *web.GetBookSheetRequest) wrapper.Result
	CreateBookByGrpc(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result
}

func NewUsecaseImpl(config *config.Configurations, logger *logger.Logger, repository repositories.Repository, gsheet googlesheet.GoogleSheetServiceInterface, grpcServices utils_grpc.GrpcServicesInterface) Usecases {
	return &UsecaseImpl{
		Config:       config,
		Logger:       logger,
		Repository:   repository,
		Gsheet:       gsheet,
		GrpcServices: grpcServices,
	}
}
