package main

import (
	"golang-api/db"
	"golang-api/utils"
	"net/http"

	books "golang-api/modules/books/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := utils.Newlogger()
	db := db.InitPostgres(logger)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		logger.Info("This service is running properly")
		return c.String(http.StatusOK, "This service is running properly")
	})

	booksGroup := e.Group("/books")
	books.New(logger, db).Mount(booksGroup)

	e.Logger.Fatal(e.Start(":1234"))
}
