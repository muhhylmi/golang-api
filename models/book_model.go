package models

import (
	"golang-api/db"
	"net/http"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func GetBook() (Response, error) {
	var obj Book
	var arrObj []Book
	var res Response

	conn := db.GetConn()

	sqlString := "SELECT * FROM book"

	rows, err := conn.Query(sqlString)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Author, &obj.Year)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func CreateBook(title string, author string, year int) (Response, error) {
	var res Response
	conn := db.GetConn()

	sqlString := "INSERT book (title, author, year) VALUES (?, ?, ?)"

	stmt, err := conn.Prepare(sqlString)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, author, year)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateBook(id int, title string, author string, year int) (Response, error) {
	var res Response
	conn := db.GetConn()

	sqlString := "SELECT * FROM book WHERE id=?"
	_, err := conn.Exec(sqlString, id)
	if err != nil {
		return res, err
	}

	sqlStringUpdate := "UPDATE book set title=?, author=?, year=? where id=?"
	stmt, err := conn.Prepare(sqlStringUpdate)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, author, year, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, err
}

func DeleteBook(id int) (Response, error) {
	var res Response

	conn := db.GetConn()
	sqlString := "DELETE FROM book where id=?"
	result, err := conn.Exec(sqlString, id)
	if err != nil {
		return res, err
	}

	deletedId, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{
		"rows_affected": int(deletedId),
	}

	return res, err
}

func GetDetailBook(id int) (Response, error) {
	var res Response
	var obj Book

	conn := db.GetConn()
	sqlString := "Select * From book where id=?"
	err := conn.QueryRow(sqlString, id).Scan(&obj.Id, &obj.Title, &obj.Author, &obj.Year)
	if err != nil {
		return res, err
	}

	res.Data = map[string]interface{}{
		"title":  &obj.Title,
		"author": &obj.Author,
		"year":   &obj.Year,
	}
	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}
