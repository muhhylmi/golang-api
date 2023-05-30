package repositories

import "golang-api/modules/books/models/domain"

type Repository interface {
	Save(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	Delete(bookId uint) (uint, error)
	FindById(bookId uint) (*domain.Book, error)
	FindAll() ([]*domain.Book, error)
}
