package repository

import (
	"barterbooks/model"

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

func (r *BookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	err := r.conn.Select(&books, "SELECT * FROM books")
	return books, err
}

func (r *BookRepository) UpdateTitle(book model.Book) error {
	_, err := r.conn.NamedExec("UPDATE books SET title = :title WHERE id = :id", book)
	return err
}

func (r *BookRepository) Delete(id int) error {
	_, err := r.conn.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
