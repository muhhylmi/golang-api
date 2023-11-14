package middlewares

import (
	"golang-api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		config := config.GetConfig()
		if username == config.BASIC_AUTH_USERNAME && password == config.BASIC_AUTH_PASSWORD {
			return true, nil
		}

		return false, nil
	})
}
