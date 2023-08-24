package handler

import (
	models "golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	"golang-api/modules/books/usecases"
	"golang-api/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const contextName = "modules.books.handler"

// HTTPHandler struct
type HTTPHandler struct {
	logger  *logrus.Logger
	usecase usecases.Usecases
}

// New initiation
func New(logger *logrus.Logger, db *gorm.DB) *HTTPHandler {
	repository := repositories.NewRepositoryImpl(logger, db)
	usecaseImpl := usecases.NewUsecaseImpl(logger, repository)
	return &HTTPHandler{
		logger:  logger,
		usecase: usecaseImpl,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.GET("", h.GetAllBook)
	echoGroup.POST("", h.CreateBook)
	echoGroup.PUT("/:id", h.UpdateBook)
	echoGroup.DELETE("/:id", h.DeleteBook)
	echoGroup.GET("/:id", h.GetDetailBook)
}

func (h *HTTPHandler) GetAllBook(c echo.Context) error {
	result := h.usecase.GetBook(c.Request().Context())
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) CreateBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	book := new(models.RequestCreateBook)
	if err := utils.BindValidate(c, book); err != nil {
		log.Error("validation on request create book is failed")
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.CreateBook(c.Request().Context(), book)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) UpdateBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	book := new(models.RequestUpdateBook)
	if err := c.Bind(book); err != nil {
		log.Error("validation on request update book is failed")
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.UpdateBook(c.Request().Context(), book)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) DeleteBook(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
	id := new(models.RequestDeleteBook)
	if err := c.Bind(id); err != nil {
		log.Error("validation on request update book is failed")
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.DeleteBook(c.Request().Context(), id)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) GetDetailBook(c echo.Context) error {
	id := new(models.RequestDetailBook)
	if err := c.Bind(id); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.GetDetailBook(c.Request().Context(), id)
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}
