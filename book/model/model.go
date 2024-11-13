package model

import (
	"time"
)

type Book struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title" validate:"required,min=1,max=100"`
	Author    string    `json:"author" db:"author" validate:"required,min=1,max=100"`
	ISBN      string    `json:"isbn" db:"isbn" validate:"required,min=1,max=100"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
