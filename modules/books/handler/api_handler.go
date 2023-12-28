package handler

import (
	"golang-api/middlewares"
	models "golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	userRepo "golang-api/modules/users/repositories"

	"golang-api/modules/books/usecases"
	"golang-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const contextName = "modules.books.handler"

// HTTPHandler struct
type HTTPHandler struct {
	logger         *logrus.Logger
	userRepository userRepo.Repository
	usecase        usecases.Usecases
}

// New initiation
func New(logger *logrus.Logger, db *gorm.DB) *HTTPHandler {
	userRepo := userRepo.NewRepositoryImpl(logger, db)
	repository := repositories.NewRepositoryImpl(logger, db)
	usecaseImpl := usecases.NewUsecaseImpl(logger, repository)
	return &HTTPHandler{
		logger:         logger,
		userRepository: userRepo,
		usecase:        usecaseImpl,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.GET("", h.GetAllBook, middlewares.VerifyBearer(h.logger, h.userRepository))
	echoGroup.POST("", h.CreateBook, middlewares.VerifyBearer(h.logger, h.userRepository))
	echoGroup.PUT("/:id", h.UpdateBook, middlewares.VerifyBearer(h.logger, h.userRepository))
	echoGroup.DELETE("/:id", h.DeleteBook, middlewares.VerifyBearer(h.logger, h.userRepository))
	echoGroup.GET("/:id", h.GetDetailBook, middlewares.VerifyBearer(h.logger, h.userRepository))
	echoGroup.GET("/sheet", h.GetBookSheetData, middlewares.VerifyBearer(h.logger, h.userRepository))
}

func (h *HTTPHandler) GetAllBook(c echo.Context) error {
	result := h.usecase.GetBook(c.Request().Context())
	if result.Error != nil {
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) CreateBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	book := new(models.RequestCreateBook)
	book.Token = c.Get("user").(utils.ClaimToken)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error(err)
		perr := utils.ResultFailed(utils.NewBadRequest(err.Error()), utils.ValidationError)
		return utils.ResponseError(perr.Error, perr.StatusCode, c)
	}

	result := h.usecase.CreateBook(c.Request().Context(), book)
	if result.Error != nil {
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}

func (h *HTTPHandler) UpdateBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	book := new(models.RequestUpdateBook)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error(err)
		perr := utils.ResultFailed(utils.NewBadRequest(err.Error()), utils.ValidationError)
		return utils.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.usecase.UpdateBook(c.Request().Context(), book)
	if result.Error != nil {
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) DeleteBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	id := new(models.RequestDeleteBook)
	if err := utils.BindValidate(c, id); err != nil {
		log.Error(err)
		perr := utils.ResultFailed(utils.NewBadRequest(err.Error()), utils.ValidationError)
		return utils.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.usecase.DeleteBook(c.Request().Context(), id)
	if result.Error != nil {
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) GetDetailBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "GetDetailBook")
	id := new(models.RequestDetailBook)
	if err := utils.BindValidate(c, id); err != nil {
		log.Error(err)
		perr := utils.ResultFailed(utils.NewBadRequest(err.Error()), utils.ValidationError)
		return utils.ResponseError(perr.Error, perr.StatusCode, c)
	}
	result := h.usecase.GetDetailBook(c.Request().Context(), id)
	if result.Error != nil {
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) GetBookSheetData(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "GetBookSheetData")
	result := h.usecase.GetBookSheetData(c.Request().Context())
	if result.Error != nil {
		log.Error(result.Error)
		return utils.ResponseError(result.Error, result.StatusCode, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}
