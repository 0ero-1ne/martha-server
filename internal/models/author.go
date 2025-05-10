package models

import "time"

type Author struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fullname  string    `json:"fullname" binding:"required"`
	Biography string    `json:"biography" binding:"required"`
	Image     string    `json:"image"`
	Books     []*Book   `json:"books,omitempty" gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
