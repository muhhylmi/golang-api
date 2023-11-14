package main

import (
	"fmt"
	"golang-api/config"
	"golang-api/db"
	"golang-api/utils"
	"net/http"

	books "golang-api/modules/books/handler"
	users "golang-api/modules/users/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	logger := utils.Newlogger()
	db := db.InitPostgres(logger)
	e := echo.New()
	e.Validator = utils.NewValidationUtil()

	e.GET("/", func(c echo.Context) error {
		logger.Info("This service is running properly")
		return c.String(http.StatusOK, "This service is running properly")
	})

	booksGroup := e.Group("/books")
	userGroup := e.Group("/users")
	books.New(logger, db).Mount(booksGroup)
	users.New(logger, db).Mount(userGroup)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.HOST)))
}
