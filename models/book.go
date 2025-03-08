package models

import "time"

type Book struct {
	Id          uint       `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"upated_at"`
	Title       string     `json:"title" gorm:"not null;default:null"`
	Description string     `json:"description" gorm:"not null;default:null"`
	Status      string     `json:"status" gorm:"not null;default:null"`
	Year        int        `json:"year" gorm:"not null;default:null"`
	Views       int        `json:"views" gorm:"not null;default:null"`
	Cover       string     `json:"image"`
	Authors     []*Author  `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"authors"`
	Tags        []*Tag     `gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"tags"`
	Comments    []Comment  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments"`
	BookRates   []BookRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
	Chapters    []Chapter  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"chapters"`
}
