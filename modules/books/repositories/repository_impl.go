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

func NewRepositoryImpl(logger *logrus.Logger, db *gorm.DB) Repository {
	return &RepositoryImpl{
		logger: logger,
		db:     db,
	}
}

func (repository *RepositoryImpl) FindAll() ([]*domain.Book, error) {
	var books []*domain.Book
	result := repository.db.Model(&domain.Book{}).Find(&books)
	return books, result.Error
}

func (repository *RepositoryImpl) Save(book *domain.Book) (*domain.Book, error) {
	result := repository.db.Create(&book)
	return book, result.Error
}

func (repository *RepositoryImpl) Update(book *domain.Book) (*domain.Book, error) {
	result := repository.db.Save(&book)
	return book, result.Error
}

func (repository *RepositoryImpl) Delete(id uint) (uint, error) {
	result := repository.db.Delete(&domain.Book{}, id)
	return id, result.Error
}

func (repository *RepositoryImpl) FindById(id uint) (*domain.Book, error) {
	var book *domain.Book
	result := repository.db.Where(&domain.Book{Id: id}).First(&book)
	return book, result.Error
}
