package repository

import (
	"user/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) Create(user model.User) error {
	_, err := r.conn.NamedExec("INSERT INTO users (username, first_name, last_name, role) VALUES (:username, :first_name, :last_name, :role)", user)
	return err
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.conn.Select(&users, "SELECT * FROM users")
	return users, err
}

func (r *UserRepository) GetUser(id int) (model.User, error) {
	var user model.User
	err := r.conn.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	return user, err
}

func (r *UserRepository) UpdateUser(user model.User, id int) error {
	_, err := r.conn.Exec("UPDATE users SET title = $2, author = $3, isbn = $4 WHERE id = $1",
		id,
		user.Title,
		user.Author,
		user.ISBN,
	)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.conn.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
