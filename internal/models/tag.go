package models

import (
	"time"
)

type Tag struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title" binding:"required,min=2"`
	Books     []*Book   `json:"books,omitempty" gorm:"many2many:books_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
