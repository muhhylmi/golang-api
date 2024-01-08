package usecases

import (
	"context"
	"fmt"
	"golang-api/modules/books/models/domain"
	"golang-api/modules/books/models/web"
	"golang-api/utils/constant"
	"golang-api/utils/wrapper"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (usecase *UsecaseImpl) GetBook(ctx context.Context) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "GetBook")
	categories, err := usecase.Repository.FindAll()
	if err != nil {
		log.Error("Book is not found")
		return wrapper.ResultFailed(wrapper.NewNotFound("Book Is Not Found"), constant.BookNotFound)
	}
	return wrapper.ResultSuccess(categories)
}

func (usecase *UsecaseImpl) CreateBook(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "CreateBook")
	bookData := domain.Book{
		Id:        uuid.New().String(),
		Title:     payload.Title,
		Author:    payload.Author,
		Year:      payload.Year,
		Price:     payload.Price,
		CreatedBy: payload.Token.UserId,
		CreatedAt: time.Now().Unix(),
	}
	book, err := usecase.Repository.Save(&bookData)
	if err != nil {
		log.Error(err.Error())
		return wrapper.ResultFailed(wrapper.NewBadRequest("Cannot Create Book"), constant.CreateBookFailed)
	}
	return wrapper.ResultSuccess(book)
}

func (usecase *UsecaseImpl) DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) wrapper.Result {
	book, err := usecase.Repository.Delete(payload.Id)
	if err != nil {
		return wrapper.ResultFailed(wrapper.NewBadRequest("Cannot delete book"), constant.DeleteBookFailed)
	}
	return wrapper.ResultSuccess(book)
}

func (usecase *UsecaseImpl) UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) wrapper.Result {
	bookData := domain.Book{
		Id:        payload.Id,
		Title:     payload.Title,
		Price:     payload.Price,
		Author:    payload.Author,
		Year:      payload.Year,
		UpdatedBy: payload.Token.UserId,
	}
	book, err := usecase.Repository.Update(&bookData)
	if err != nil {
		return wrapper.ResultFailed(wrapper.NewBadRequest("cannot update book"), constant.UpdateBookFailed)
	}
	return wrapper.ResultSuccess(book)
}

func (usecase *UsecaseImpl) GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) wrapper.Result {
	book, err := usecase.Repository.FindById(payload.Id)
	if err != nil {
		return wrapper.ResultFailed(wrapper.NewNotFound("Books Is not Found"), constant.BookNotFound)
	}
	return wrapper.ResultSuccess(book)
}

func (usecase *UsecaseImpl) GetBookSheetData(ctx context.Context, payload *web.GetBookSheetRequest) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "GetBookSheetData")

	values, filteredRow, err := usecase.Gsheet.GetSheetDataWithFilter(usecase.Config.SPREAD_SHEET_ID, "Sheet2!A2:E")
	if err != nil {
		log.Error(err)
		error := wrapper.NewBadRequest("failed to get sheet data")
		return wrapper.ResultFailed(error, constant.FailedGetBookSheet)
	}

	var resultGsheet [][]interface{}
	gsheetValue := values.Values
	if filteredRow != nil || len(filteredRow) > 0 {
		for _, idx := range filteredRow {
			if idx >= 0 && idx < len(gsheetValue) {
				resultGsheet = append(resultGsheet, gsheetValue[idx])
			}
		}
		gsheetValue = resultGsheet
	}

	if payload.Author != "" || payload.Title != "" {
		gsheetValue = filterRowsByColumnValue(values.Values, payload)
	}

	responseData := make([]web.ResponseSheetBook, len(gsheetValue))
	for idx1, value := range gsheetValue {
		count := float64(0)
		price := float64(0)
		for idx, data := range value {
			if idx == 0 {
				responseData[idx1].BookName = fmt.Sprintf("%v", data)
			} else if idx == 1 {
				responseData[idx1].Author = fmt.Sprintf("%v", data)
			} else if idx == 2 {
				responseData[idx1].Year = fmt.Sprintf("%v", data)
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
		responseData[idx1].Total = count * price
	}
	return wrapper.ResultSuccess(responseData)
}

func filterRowsByColumnValue(rows [][]interface{}, req *web.GetBookSheetRequest) [][]interface{} {
	var filteredRows [][]interface{}

	for _, row := range rows {
		regexAuthor := regexp.MustCompile(strings.ToLower(req.Author))
		authorValue := strings.ToLower(fmt.Sprintf("%v", row[1]))
		regexTitle := regexp.MustCompile(strings.ToLower(req.Title))
		titleValue := strings.ToLower(fmt.Sprintf("%v", row[0]))

		if regexAuthor.MatchString(authorValue) && regexTitle.MatchString(titleValue) {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

func (usecase *UsecaseImpl) CreateBookByGrpc(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result {
	log := usecase.Logger.LogWithContext(contextName, "CreateBookByGrpc")
	bookData := domain.Book{
		Id:        uuid.New().String(),
		Title:     payload.Title,
		Author:    payload.Author,
		Year:      payload.Year,
		Price:     payload.Price,
		CreatedBy: payload.Token.UserId,
		CreatedAt: time.Now().Unix(),
	}
	book, err := usecase.Repository.Save(&bookData)
	if err != nil {
		log.Error(err.Error())
		return wrapper.ResultFailed(wrapper.NewBadRequest("Cannot Create Book"), constant.CreateBookFailed)
	}
	return wrapper.ResultSuccess(book)
}
