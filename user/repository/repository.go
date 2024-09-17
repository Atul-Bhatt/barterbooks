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
	_, err := r.conn.NamedExec("INSERT INTO users (username, first_name, last_name, user_role) VALUES (:username, :firstname, :lastname, :role)", user)
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
	_, err := r.conn.Exec("UPDATE users SET username = $2, first_name = $3, last_name = $4 user_role = $5 WHERE id = $1",
		id,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Role,
	)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.conn.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
