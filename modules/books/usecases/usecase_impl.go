package usecases

import (
	"context"
	"golang-api/modules/books/models/domain"
	"golang-api/modules/books/models/web"
	"golang-api/modules/books/repositories"
	"golang-api/utils"
)

type UsecaseImpl struct {
	repository repositories.Repository
}

func NewUsecaseImpl(repository repositories.Repository) Usecases {
	return &UsecaseImpl{
		repository: repository,
	}
}

func (usecase *UsecaseImpl) GetBook(ctx context.Context) utils.Result {
	var result utils.Result
	categories, err := usecase.repository.FindAll()
	if err != nil {
		error := utils.NewBadRequest()
		result.Error = error
		return result
	}
	result.Data = categories
	return result
}

func (usecase *UsecaseImpl) CreateBook(ctx context.Context, payload *web.RequestCreateBook) utils.Result {
	var result utils.Result
	bookData := domain.Book{
		Title:  payload.Title,
		Author: payload.Author,
		Year:   payload.Year,
	}
	book, err := usecase.repository.Save(bookData)
	if err != nil {
		error := utils.NewBadRequest()
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
		error := utils.NewBadRequest()
		result.Error = error
		return result
	}
	result.Data = book
	return result
}

func (usecase *UsecaseImpl) UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) utils.Result {
	var result utils.Result
	bookData := domain.Book{
		Id:     payload.Id,
		Title:  payload.Title,
		Author: payload.Author,
		Year:   payload.Year,
	}
	book, err := usecase.repository.Update(bookData)
	if err != nil {
		error := utils.NewBadRequest()
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
		error := utils.NewBadRequest()
		result.Error = error
		return result
	}
	result.Data = book
	return result
}
