package handler

import (
	"context"
	"fmt"
	models "golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	userRepo "golang-api/modules/users/repositories"
	"golang-api/proto"
	"golang-api/utils/app"
	"golang-api/utils/config"
	"golang-api/utils/constant"
	"golang-api/utils/jwt"
	"golang-api/utils/logger"
	"golang-api/utils/middlewares"
	utils "golang-api/utils/validator"
	"golang-api/utils/wrapper"

	"golang-api/modules/books/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const contextName = "modules.books.handler"

// HTTPHandler struct
type HTTPHandler struct {
	proto.UnimplementedBookServiceServer

	Logger         *logger.Logger
	UserRepository userRepo.Repository
	Usecase        usecases.Usecases
	Validator      echo.Validator
	Config         *config.Configurations

	GrpcServer *grpc.Server
}

// New initiation
func New(apps *app.App) *HTTPHandler {
	userRepo := userRepo.NewRepositoryImpl(apps.Logger, apps.DBService)
	repository := repositories.NewRepositoryImpl(apps.Logger, apps.DBService)
	usecaseImpl := usecases.NewUsecaseImpl(apps.GlobalConfig, apps.Logger, repository, apps.GsheetService, apps.GrpcServices)
	return &HTTPHandler{
		Logger:         apps.Logger,
		UserRepository: userRepo,
		Usecase:        usecaseImpl,
		Validator:      apps.Validator,
		Config:         apps.GlobalConfig,
		GrpcServer:     apps.GRPC,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.GET("", h.GetAllBook, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.POST("", h.CreateBook, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.PUT("/:id", h.UpdateBook, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.DELETE("/:id", h.DeleteBook, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.GET("/:id", h.GetDetailBook, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.GET("/sheet", h.GetBookSheetData, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))
	echoGroup.POST("/rpc", h.CreateBookByGrpc, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepository))

	h.GrpcServer.RegisterService(&proto.BookService_ServiceDesc, h)
}

func (h *HTTPHandler) GetAllBook(c echo.Context) error {
	result := h.Usecase.GetBook(c.Request().Context())
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) CreateBook(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateBook")
	book := new(models.RequestCreateBook)
	book.Token = c.Get("user").(jwt.ClaimToken)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}

	result := h.Usecase.CreateBook(c.Request().Context(), book)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}

func (h *HTTPHandler) UpdateBook(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateBook")
	book := new(models.RequestUpdateBook)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.Usecase.UpdateBook(c.Request().Context(), book)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}
	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) DeleteBook(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateBook")
	id := new(models.RequestDeleteBook)
	if err := utils.BindValidate(c, id); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.Usecase.DeleteBook(c.Request().Context(), id)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}
	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) GetDetailBook(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "GetDetailBook")
	id := new(models.RequestDetailBook)
	if err := utils.BindValidate(c, id); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.Usecase.GetDetailBook(c.Request().Context(), id)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}
	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) GetBookSheetData(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "GetBookSheetData")
	model := new(models.GetBookSheetRequest)
	if err := utils.BindValidate(c, model); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.Usecase.GetBookSheetData(c.Request().Context(), model)
	if result.Error != nil {
		log.Error(result.Error)
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}
	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) CreateBookByGrpc(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateBookByGrpc")
	book := new(models.RequestCreateBook)
	book.Token = c.Get("user").(jwt.ClaimToken)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return wrapper.ResponseError(perr.Error, perr.StatusCode, c)
	}

	result := h.Usecase.CreateBookByGrpc(c.Request().Context(), book)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}

func (h *HTTPHandler) GrpcCreateBook(c context.Context, req *proto.BookDataRequest) (*proto.BookDataResponse, error) {
	log := h.Logger.LogWithContext(contextName, "GrpcCreateBook")

	model := &models.RequestCreateBook{
		Title:  req.Name,
		Author: req.Author,
		Year:   req.Year,
		Price:  float64(req.Price),
	}
	model.Token = jwt.ClaimToken{
		UserId: "grpc",
	}

	if err := h.Validator.Validate(model); err != nil {
		log.Error(err)
		perr := wrapper.ResultFailed(wrapper.NewBadRequest(err.Error()), constant.ValidationError)
		return &proto.BookDataResponse{
			Success: false,
		}, fmt.Errorf("error validation with code " + perr.StatusCode)
	}

	result := h.Usecase.CreateBook(c, model)
	if result.Error != nil {
		return &proto.BookDataResponse{
			Success: false,
		}, fmt.Errorf("error creating book")
	}

	return &proto.BookDataResponse{
		Success: true,
	}, nil
}
