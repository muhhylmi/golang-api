package repositories

import (
	"golang-api/modules/books/models/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	logger *logrus.Logger
	db     *gorm.DB
}
type Repository interface {
	Save(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	Delete(bookId string) (string, error)
	FindById(bookId string) (*domain.Book, error)
	FindAll() ([]*domain.Book, error)
}

func NewRepositoryImpl(logger *logrus.Logger, db *gorm.DB) Repository {
	return &RepositoryImpl{
		logger: logger,
		db:     db,
	}
}
