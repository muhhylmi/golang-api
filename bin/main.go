package main

import (
	"fmt"
	"golang-api/utils/config"
	"net/http"

	books "golang-api/modules/books/handler"
	carts "golang-api/modules/cart/handler"
	users "golang-api/modules/users/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	apps := Init()
	apps.Apps.Validator = apps.Validator
	apps.Apps.GET("/", func(c echo.Context) error {
		apps.Logger.Logger.Info("This service is running properly")
		return c.String(http.StatusOK, "This service is running properly")
	})

	booksGroup := apps.Apps.Group("/books")
	userGroup := apps.Apps.Group("/users")
	cartGroup := apps.Apps.Group("/cart")
	books.New(apps).Mount(booksGroup)
	users.New(apps).Mount(userGroup)
	carts.New(apps).Mount(cartGroup)

	// run echo server
	apps.Apps.Logger.Fatal(apps.Apps.Start(fmt.Sprintf(":%s", config.GetConfig().HOST)))
}
