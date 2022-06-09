package database

import (
	"time"
)

type Author struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID        int       `json:"id" gorm:"primary_key"`
	AuthorID  uint      `json:"author_id"`
	Author    Author    `json:"author" gorm:"foreignKey:author_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
