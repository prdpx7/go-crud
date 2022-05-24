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

// type Blog struct {
// 	ID        int
// 	Author    Author
// 	Title     string
// 	Post      string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
