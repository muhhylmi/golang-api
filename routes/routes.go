package routes

import (
	"golang-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service Running Properly")
	})

	e.GET("/books", controllers.GetAllBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("books/:id", controllers.DeleteBook)
	e.GET("books/:id", controllers.GetDetailBook)

	return e
}
