package models

import "time"

type Book struct {
	Id          uint       `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"upated_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Year        int        `json:"year"`
	Views       int        `json:"views"`
	Cover       string     `json:"image"`
	Authors     []*Author  `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"authors"`
	Tags        []*Tag     `gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"tags"`
	Comments    []Comment  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments"`
	BookRates   []BookRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
	Chapters    []Chapter  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"chapters"`
}
