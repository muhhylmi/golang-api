package handler

import (
	"golang-api/middlewares"
	models "golang-api/modules/cart/models/web"
	"golang-api/modules/cart/repositories"
	userRepo "golang-api/modules/users/repositories"

	"golang-api/modules/cart/usecases"
	"golang-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const contextName = "modules.cart.handler"

// HTTPHandler struct
type HTTPHandler struct {
	logger   *logrus.Logger
	usecase  usecases.Usecases
	userRepo userRepo.Repository
}

// New initiation
func New(logger *logrus.Logger, db *gorm.DB) *HTTPHandler {
	userRepo := userRepo.NewRepositoryImpl(logger, db)
	repository := repositories.NewRepositoryImpl(logger, db)
	usecaseImpl := usecases.NewUsecaseImpl(logger, repository)
	return &HTTPHandler{
		logger:   logger,
		userRepo: userRepo,
		usecase:  usecaseImpl,
	}
}

func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.POST("", h.CreateCart, middlewares.VerifyBearer(h.logger, h.userRepo))
}

func (h *HTTPHandler) CreateCart(c echo.Context) error {
	log := utils.LogWithContext(h.logger, contextName, "CreateCart")
	cart := new(models.RequestCreateCart)
	cart.Token = c.Get("user").(utils.ClaimToken)
	if err := utils.BindValidate(c, cart); err != nil {
		log.Error(err)
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.usecase.CreateCart(c.Request().Context(), cart)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
}
