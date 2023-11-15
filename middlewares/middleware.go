package middlewares

import (
	"golang-api/config"
	"golang-api/db"
	userRepo "golang-api/modules/users/repositories"
	"golang-api/utils"

	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const contextName = "middleware"

func VerifyBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		config := config.GetConfig()
		if username == config.BASIC_AUTH_USERNAME && password == config.BASIC_AUTH_PASSWORD {
			return true, nil
		}

		return false, nil
	})
}

func VerifyBearer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := utils.LogWithContext(logrus.New(), contextName, "VerifyBearer")
			tokenString := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ")

			if len(tokenString) == 0 {
				return utils.Response(nil, "invalid token!", http.StatusUnauthorized, c)
			}
			token, err := utils.ValidateJwt(tokenString)
			if err != nil {
				return utils.Response(nil, err.Error(), http.StatusUnauthorized, c)
			}
			checkUser, err := userRepo.NewRepositoryImpl(log.Logger, db.InitPostgres(log.Logger)).FindById(token.UserId)
			if err != nil {
				return utils.Response(nil, "invalid user!", http.StatusUnauthorized, c)
			}
			claimToken := utils.ClaimToken{
				UserId:   checkUser.Id,
				Username: checkUser.Username,
				Gender:   checkUser.Gender,
			}

			c.Set("user", claimToken)
			return next(c)
		}
	}
}
