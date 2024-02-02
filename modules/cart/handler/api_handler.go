package handler

import (
	models "golang-api/modules/cart/models/web"
	"golang-api/modules/cart/repositories"
	"golang-api/modules/cart/usecases"
	userRepo "golang-api/modules/users/repositories"
	app_type "golang-api/utils/app_opts"
	"golang-api/utils/config"
	"golang-api/utils/jwt"
	"golang-api/utils/logger"
	"golang-api/utils/middlewares"
	utils "golang-api/utils/validator"
	"golang-api/utils/wrapper"
	"net/http"

	"github.com/labstack/echo/v4"
)

const contextName = "modules.cart.handler"

// HTTPHandler struct
type HTTPHandler struct {
	Logger   *logger.Logger
	UseCase  usecases.Usecases
	UserRepo userRepo.Repository
	Config   *config.Configurations
}

// New initiation
func New(apps *app_type.App) *HTTPHandler {
	userRepo := userRepo.NewRepositoryImpl(apps.Logger, apps.DBService)
	repository := repositories.NewRepositoryImpl(apps.Logger, apps.DBService)
	usecaseImpl := usecases.NewUsecaseImpl(apps.GlobalConfig, apps.Logger, repository)
	return &HTTPHandler{
		Logger:   apps.Logger,
		UserRepo: userRepo,
		UseCase:  usecaseImpl,
		Config:   apps.GlobalConfig,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.POST("", h.CreateCart, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepo))
	echoGroup.GET("", h.GetAllCart, middlewares.VerifyBearer(h.Logger, h.Config, h.UserRepo))
}

func (h *HTTPHandler) CreateCart(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "CreateCart")
	cart := new(models.RequestCreateCart)
	cart.Token = c.Get("user").(jwt.ClaimToken)
	if err := utils.BindValidate(c, cart); err != nil {
		log.Error(err)
		return wrapper.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.UseCase.CreateCart(c.Request().Context(), cart)
	if result.Error != nil {
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}

func (h *HTTPHandler) GetAllCart(c echo.Context) error {
	log := h.Logger.LogWithContext(contextName, "GetAllCart")
	cart := new(models.RequestListCart)
	if err := utils.BindValidate(c, cart); err != nil {
		log.Error(err)
		return wrapper.Response(nil, err.Error(), http.StatusBadRequest, c)
	}
	result := h.UseCase.GetAllCart(c.Request().Context(), cart)
	if result.Error != nil {
		log.Error(result.Error)
		return wrapper.ResponseError(result.Error, result.StatusCode, c)
	}

	return wrapper.Response(result.Data, "Your Request has been Approve", http.StatusOK, c)
}
