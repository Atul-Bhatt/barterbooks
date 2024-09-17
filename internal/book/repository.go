package repository

import (
	"book/model"

	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	conn *sqlx.DB
}

func NewBookRepository(conn *sqlx.DB) *BookRepository {
	return &BookRepository{conn: conn}
}

func (r *BookRepository) Create(book model.Book) error {
	_, err := r.conn.NamedExec("INSERT INTO books (title, author, isbn) VALUES (:title, :author, :isbn)", book)
	return err
}

func (r *BookRepository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.conn.Select(&books, "SELECT * FROM books")
	return books, err
}

func (r *BookRepository) GetBook(id int) (model.Book, error) {
	var book model.Book
	err := r.conn.Get(&book, "SELECT * FROM books WHERE id = $1", id)

	return book, err
}

func (r *BookRepository) UpdateBook(book model.Book, id int) error {
	_, err := r.conn.Exec("UPDATE books SET title = $2, author = $3, isbn = $4 WHERE id = $1",
		id,
		book.Title,
		book.Author,
		book.ISBN,
	)
	return err
}

func (r *BookRepository) DeleteBook(id int) error {
	_, err := r.conn.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
