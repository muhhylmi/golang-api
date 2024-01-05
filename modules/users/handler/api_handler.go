package handler

import (
	models "golang-api/modules/users/models/web"
	"golang-api/modules/users/repositories"
	"golang-api/modules/users/usecases"
	"golang-api/utils"
	"golang-api/utils/app"
	"golang-api/utils/logger"
	"golang-api/utils/middlewares"
	"golang-api/utils/wrapper"
	"net/http"

	"github.com/labstack/echo/v4"
)

const contextName = "modules.users.handler"

// HTTPHandler struct
type HTTPHandler struct {
	Logger  *logger.Logger
	UseCase usecases.Usecases
}

// New initiation
func New(apps *app.App) *HTTPHandler {
	repository := repositories.NewRepositoryImpl(apps.Logger, apps.DBService)
	usecaseImpl := usecases.NewUsecaseImpl(apps.GlobalConfig, apps.Logger, repository)
	return &HTTPHandler{
		Logger:  apps.Logger,
		UseCase: usecaseImpl,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.POST("", h.CreateUser, middlewares.VerifyBasicAuth())
	echoGroup.POST("/login", h.LoginUser, middlewares.VerifyBasicAuth())
}

func (h *HTTPHandler) CreateUser(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateUser")
	user := new(models.RequestCreateUser)
	if err := utils.BindValidate(c, user); err != nil {
		log.Error(err)
		return wrapper.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.UseCase.CreateUser(c.Request().Context(), user)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}

func (h *HTTPHandler) LoginUser(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "LoginUser")
	user := new(models.RequestLogin)
	if err := utils.BindValidate(c, user); err != nil {
		log.Error(err)
		return wrapper.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.UseCase.LoginUser(c.Request().Context(), user)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}
