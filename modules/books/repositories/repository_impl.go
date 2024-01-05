package repositories

import (
	"golang-api/modules/books/models/domain"
)

func (repository *RepositoryImpl) FindAll() ([]*domain.Book, error) {
	var books []*domain.Book
	result := repository.DB.Gorm.Model(&domain.Book{}).Find(&books)
	return books, result.Error
}

func (repository *RepositoryImpl) Save(book *domain.Book) (*domain.Book, error) {
	result := repository.DB.Gorm.Create(&book)
	return book, result.Error
}

func (repository *RepositoryImpl) Update(book *domain.Book) (*domain.Book, error) {
	result := repository.DB.Gorm.Save(&book)
	return book, result.Error
}

func (repository *RepositoryImpl) Delete(id string) (string, error) {
	result := repository.DB.Gorm.Delete(&domain.Book{Id: id})
	return id, result.Error
}

func (repository *RepositoryImpl) FindById(id string) (*domain.Book, error) {
	var book *domain.Book
	result := repository.DB.Gorm.Where(&domain.Book{Id: id}).First(&book)
	return book, result.Error
}
