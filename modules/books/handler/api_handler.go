package handler

import (
	"fmt"
	models "golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	"golang-api/modules/books/usecases"
	"golang-api/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// HTTPHandler struct
type HTTPHandler struct {
	usecase usecases.Usecases
}

// New initiation
func New(db *gorm.DB) *HTTPHandler {
	repository := repositories.NewRepositoryImpl(db)
	usecaseImpl := usecases.NewUsecaseImpl(repository)
	return &HTTPHandler{
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
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) CreateBook(c echo.Context) error {
	book := new(models.RequestCreateBook)
	fmt.Println(book)
	if err := c.Bind(book); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.CreateBook(c.Request().Context(), book)
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) UpdateBook(c echo.Context) error {
	book := new(models.RequestUpdateBook)
	if err := c.Bind(book); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.UpdateBook(c.Request().Context(), book)
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}

func (h *HTTPHandler) DeleteBook(c echo.Context) error {
	id := new(models.RequestDeleteBook)
	if err := c.Bind(id); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.usecase.DeleteBook(c.Request().Context(), id)
	if result.Error != nil {
		log.Println(result.Error)
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
