package logic

import (
	"book/graph/database"
	"book/graph/model"
)

func CreateNewBook(input model.BookInput) (*model.Book, error) {
	if input.BookName != "" && input.BookDescription != "" && input.BookAuthorName != "" {
		query := `INSERT INTO newbook(name, description,author_name) VALUES($1,$2,$3)`
		db := database.GetDB()
		_, err := db.Exec(query, input.BookName, input.BookDescription, input.BookAuthorName)
		if err != nil {
			return nil, err
		} else {
			res := &model.Book{
				BookName:        &input.BookName,
				BookDescription: &input.BookDescription,
				BookAuthorName:  &input.BookAuthorName}
			return res, nil
		}
	}
	return nil, nil
}

func GetBooks() ([]*model.Book, error) {
	var book model.Book
	var books []*model.Book
	Query := `SELECT name, description, author_name FROM books`
	db := database.GetDB()
	rows, err := db.Query(Query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(book.ID, book.BookName, book.BookDescription, book.BookAuthorName)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func GetBook(input model.BookID) (*model.Book, error) {
	var book model.Book
	Query := `SELECT name, description, author_name FROM books WHERE id=$1`
	db := database.GetDB()
	err := db.QueryRow(Query, input.ID).Scan(book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
