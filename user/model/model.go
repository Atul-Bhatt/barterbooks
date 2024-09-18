package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	FirstName string    `json:"firstname" db:"first_name"`
	LastName  string    `json:"lastname" db:"last_name"`
	Role      string    `json:"role" db:"user_role"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
