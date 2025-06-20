package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}