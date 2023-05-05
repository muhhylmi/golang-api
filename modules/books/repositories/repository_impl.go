package repositories

import (
	"database/sql"
	"golang-api/modules/books/models/domain"

	"github.com/labstack/gommon/log"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (repository *RepositoryImpl) FindAll() ([]domain.Book, error) {
	var book domain.Book
	var books []domain.Book
	sqlString := "SELECT * FROM book"

	rows, err := repository.db.Query(sqlString)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

func (repository *RepositoryImpl) Save(book domain.Book) (domain.Book, error) {
	sqlString := "INSERT book (title, author, year) VALUES (?, ?, ?)"
	stmt, err := repository.db.Prepare(sqlString)
	if err != nil {
		return book, err
	}

	result, err := stmt.Exec(book.Title, book.Author, book.Year)
	if err != nil {
		return book, err
	}

	bookId, err := result.LastInsertId()
	if err != nil {
		return book, err
	}

	book.Id = int(bookId)
	return book, nil
}

func (repository *RepositoryImpl) Update(book domain.Book) (domain.Book, error) {
	sqlString := "SELECT * FROM book WHERE id=?"
	_, err := repository.db.Exec(sqlString, book.Id)
	if err != nil {
		return book, err
	}

	sqlStringUpdate := "UPDATE book set title=?, author=?, year=? where id=?"
	stmt, err := repository.db.Prepare(sqlStringUpdate)
	if err != nil {
		return book, err
	}

	result, err := stmt.Exec(book.Title, book.Author, book.Year, book.Id)
	if err != nil {
		return book, err
	}
	log.Debug(result.LastInsertId())
	return book, nil
}

func (repository *RepositoryImpl) Delete(id int) (int, error) {
	sqlString := "DELETE FROM book where id=?"
	result, err := repository.db.Exec(sqlString, id)
	if err != nil {
		return 0, err
	}

	deletedId, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(deletedId), nil
}

func (repository *RepositoryImpl) FindById(id int) (domain.Book, error) {
	var book domain.Book

	sqlString := "Select * From book where id=?"
	err := repository.db.QueryRow(sqlString, id).Scan(&book.Id, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return book, err
	}
	return book, nil
}
