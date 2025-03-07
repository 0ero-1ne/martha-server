package models

import "time"

type Author struct {
	Id        uint      `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
	Fullname  string    `json:"fullname"`
	Biography string    `json:"biography"`
	Image     string    `json:"image"`
	Books     []*Book   `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"books"`
}
