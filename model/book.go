package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	ISBN      string    `json:"isbn" db:"isbn"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
