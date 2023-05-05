package controllers

import (
	"golang-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBook(c echo.Context) error {
	result, err := models.GetBook()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateBook(c echo.Context) error {
	book := new(models.RequestCreateBook)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	result, err := models.CreateBook(book.Title, book.Author, book.Year)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateBook(c echo.Context) error {
	book := new(models.RequestCreateBook)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	id := c.Param("id")
	params_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.UpdateBook(params_id, book.Title, book.Author, book.Year)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	params_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	result, err := models.DeleteBook(params_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetDetailBook(c echo.Context) error {
	id := c.Param("id")
	params_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	result, err := models.GetDetailBook(params_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
