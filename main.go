package main

import (
	"golang-api/db"
	"net/http"

	books "golang-api/modules/books/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This service is running properly")
	})

	booksGroup := e.Group("/books")
	books.New().Mount(booksGroup)

	e.Logger.Fatal(e.Start(":1234"))
}
