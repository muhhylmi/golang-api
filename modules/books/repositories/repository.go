package repositories

import "golang-api/modules/books/models/domain"

type Repository interface {
	Save(book domain.Book) (domain.Book, error)
	Update(book domain.Book) (domain.Book, error)
	Delete(bookId int) (int, error)
	FindById(bookId int) (domain.Book, error)
	FindAll() ([]domain.Book, error)
}
