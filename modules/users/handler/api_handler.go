package handler

import (
	models "golang-api/modules/users/models/web"
	"golang-api/modules/users/repositories"
	"golang-api/modules/users/usecases"
	"golang-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const contextName = "modules.users.handler"

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
	echoGroup.POST("", h.CreateUser)
}

func (h *HTTPHandler) CreateUser(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateUser")
	user := new(models.RequestCreateUser)
	if err := utils.BindValidate(c, user); err != nil {
		log.Error(err)
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.usecase.CreateUser(c.Request().Context(), user)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}
