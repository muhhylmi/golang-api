package usecases

import (
	"context"
	"fmt"
	"golang-api/config"
	"golang-api/modules/books/models/domain"
	"golang-api/modules/books/models/web"
	"golang-api/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (usecase *UsecaseImpl) GetBook(ctx context.Context) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "GetBook")
	var result utils.Result
	categories, err := usecase.repository.FindAll()
	if err != nil {
		log.Error("Book is not found")
		error := utils.NewBadRequest("Book Is Not Found")
		result.Error = error
		return result
	}
	result.Data = categories
	return result
}

func (usecase *UsecaseImpl) CreateBook(ctx context.Context, payload *web.RequestCreateBook) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateBook")
	var result utils.Result
	bookData := domain.Book{
		Id:        uuid.New().String(),
		Title:     payload.Title,
		Author:    payload.Author,
		Year:      payload.Year,
		Price:     payload.Price,
		CreatedBy: payload.Token.UserId,
		CreatedAt: time.Now().Unix(),
	}
	book, err := usecase.repository.Save(&bookData)
	if err != nil {
		log.Error(err.Error())
		error := utils.NewBadRequest("Cannot Create Book")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) utils.Result {
	var result utils.Result
	book, err := usecase.repository.Delete(payload.Id)
	if err != nil {
		error := utils.NewBadRequest("Cannot delete book")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) utils.Result {
	var result utils.Result
	bookData := domain.Book{
		Id:        payload.Id,
		Title:     payload.Title,
		Price:     payload.Price,
		Author:    payload.Author,
		Year:      payload.Year,
		UpdatedBy: payload.Token.UserId,
	}
	book, err := usecase.repository.Update(&bookData)
	if err != nil {
		error := utils.NewBadRequest("cannot update book")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) utils.Result {
	var result utils.Result
	book, err := usecase.repository.FindById(payload.Id)
	if err != nil {
		error := utils.NewNotFound("Books Is not Found")
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) GetBookSheetData(ctx context.Context) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "GetBookSheetData")

	var result utils.Result
	srv, err := utils.GetSheetConfig(log)
	if err != nil {
		error := utils.NewConflict("failed to get sheet config")
		result.Error = error
		return result
	}
	values, err := utils.GetSheetData(srv, config.GetConfig().SPREAD_SHEET_ID, "Sheet1!A:E")
	if err != nil {
		log.Error(err)
		error := utils.NewBadRequest("failed to get sheet data")
		result.Error = error
		return result
	}
	responseData := make([]web.ResponseSheetBook, len(values.Values)-1)
	for idx1, value := range values.Values {
		if idx1 == 0 {
			continue
		}
		count := float64(0)
		price := float64(0)
		for idx, data := range value {
			if idx == 0 {
				responseData[idx1-1].BookName = fmt.Sprintf("%v", data)
			} else if idx == 1 {
				responseData[idx1-1].Author = fmt.Sprintf("%v", data)
			} else if idx == 2 {
				responseData[idx1-1].Year = fmt.Sprintf("%v", data)
			} else {
				if idx == 3 {
					temp := fmt.Sprintf("%v", data)
					count, _ = strconv.ParseFloat(temp, 64)
				}
				if idx == 4 {
					temp := fmt.Sprintf("%v", data)
					price, _ = strconv.ParseFloat(temp, 64)
				}
			}
		}
		responseData[idx1-1].Total = count * price
	}
	result.Data = responseData

	return result
}
