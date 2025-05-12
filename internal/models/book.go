package models

import "time"

type Book struct {
	Id          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Title       string       `json:"title" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Status      string       `json:"status" binding:"required"`
	Year        int          `json:"year" binding:"required,min=0"`
	Views       int          `json:"views"`
	Cover       string       `json:"cover"`
	Authors     []*Author    `json:"authors,omitempty"  gorm:"many2many:books_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Tags        []*Tag       `json:"tags,omitempty"     gorm:"many2many:books_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Comments    []Comment    `json:"comments,omitempty"`
	BooksRates  []BooksRates `json:"rates,omitempty"`
	Chapters    []Chapter    `json:"chapters,omitempty"`
}
