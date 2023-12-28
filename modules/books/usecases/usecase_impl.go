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
	categories, err := usecase.repository.FindAll()
	if err != nil {
		log.Error("Book is not found")
		return utils.ResultFailed(utils.NewBadRequest("Book Is Not Found"), utils.BookNotFound)
	}
	return utils.ResultSuccess(categories)
}

func (usecase *UsecaseImpl) CreateBook(ctx context.Context, payload *web.RequestCreateBook) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "CreateBook")
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
		return utils.ResultFailed(utils.NewBadRequest("Cannot Create Book"), utils.CreateBookFailed)
	}
	return utils.ResultSuccess(book)
}

func (usecase *UsecaseImpl) DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) utils.Result {
	book, err := usecase.repository.Delete(payload.Id)
	if err != nil {
		return utils.ResultFailed(utils.NewBadRequest("Cannot delete book"), utils.DeleteBookFailed)
	}
	return utils.ResultSuccess(book)
}

func (usecase *UsecaseImpl) UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) utils.Result {
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
		return utils.ResultFailed(utils.NewBadRequest("cannot update book"), utils.UpdateBookFailed)
	}
	return utils.ResultSuccess(book)
}

func (usecase *UsecaseImpl) GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) utils.Result {
	book, err := usecase.repository.FindById(payload.Id)
	if err != nil {
		return utils.ResultFailed(utils.NewNotFound("Books Is not Found"), utils.BookNotFound)
	}
	return utils.ResultSuccess(book)
}

func (usecase *UsecaseImpl) GetBookSheetData(ctx context.Context) utils.Result {
	log := utils.LogWithContext(usecase.logger, contextName, "GetBookSheetData")

	srv, err := utils.GetSheetConfig(log)
	if err != nil {
		error := utils.NewConflict("failed to get sheet config")
		return utils.ResultFailed(error, utils.FailedConnectSheet)
	}
	values, err := utils.GetSheetData(srv, config.GetConfig().SPREAD_SHEET_ID, "Sheet1!A:E")
	if err != nil {
		log.Error(err)
		error := utils.NewBadRequest("failed to get sheet data")
		return utils.ResultFailed(error, utils.FailedGetBookSheet)

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
	return utils.ResultSuccess(responseData)
}
