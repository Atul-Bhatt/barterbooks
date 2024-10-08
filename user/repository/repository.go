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
	args := []interface{}{
		user.Username,
		user.FirstName,
		user.LastName,
		user.Role,
		user.Password,
	}
	_, err := r.conn.Exec("INSERT INTO users (username, first_name, last_name, user_role, password) VALUES ($1, $2, $3, $4, $5)", args...)
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

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.conn.Get(&user, "SELECT * FROM users WHERE username = $1", username)

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

func (r *UserRepository) GetDBPassword(username string) (string, error) {
	var dbPassword string

	err := r.conn.Get(&dbPassword, "SELECT password FROM users where username = $1", username)
	if err != nil {
		return dbPassword, err
	}

	return dbPassword, nil
}

func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var count int
	err := r.conn.Get(&count, "SELECT COUNT(*) FROM users WHERE username = $1", username)

	if count > 0 {
		return true, err
	}

	return false, err
}
