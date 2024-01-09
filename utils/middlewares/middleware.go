package middlewares

import (
	"golang-api/modules/users/repositories"
	"golang-api/utils/config"
	"golang-api/utils/jwt"
	"golang-api/utils/logger"
	"golang-api/utils/wrapper"

	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const contextName = "middleware"

func VerifyBasicAuth(config *config.Configurations) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		if username == config.BASIC_AUTH_USERNAME && password == config.BASIC_AUTH_PASSWORD {
			return true, nil
		}

		return false, nil
	})
}

func VerifyBearer(logger *logger.Logger, config *config.Configurations, repository repositories.Repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := logger.LogWithContext(contextName, "VerifyBearer")
			tokenString := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ")

			if len(tokenString) == 0 {
				return wrapper.Response(nil, "invalid token!", http.StatusUnauthorized, c)
			}
			token, err := jwt.ValidateJwt(tokenString, config)
			if err != nil {
				log.Error(err.Error())
				return wrapper.Response(nil, err.Error(), http.StatusUnauthorized, c)
			}
			checkUser, err := repository.FindById(token.UserId)
			if err != nil {
				log.Error(err.Error())
				return wrapper.Response(nil, "invalid user!", http.StatusUnauthorized, c)
			}
			claimToken := jwt.ClaimToken{
				UserId:   checkUser.Id,
				Username: checkUser.Username,
				Gender:   checkUser.Gender,
			}

			c.Set("user", claimToken)
			return next(c)
		}
	}
}
