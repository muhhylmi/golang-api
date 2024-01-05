package repositories

import (
	"golang-api/modules/books/models/domain"
	"golang-api/utils/database"
	"golang-api/utils/logger"
)

type RepositoryImpl struct {
	Logger *logger.Logger
	DB     *database.DBService
}
type Repository interface {
	Save(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	Delete(bookId string) (string, error)
	FindById(bookId string) (*domain.Book, error)
	FindAll() ([]*domain.Book, error)
}

func NewRepositoryImpl(logger *logger.Logger, db *database.DBService) Repository {
	return &RepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
