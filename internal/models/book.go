package models

import "time"

type Book struct {
	Id          uint       `json:"id"                 gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"              gorm:"not null;default:null"                                               binding:"required"`
	Description string     `json:"description"        gorm:"not null;default:null"                                               binding:"required"`
	Status      string     `json:"status"             gorm:"not null;default:null"                                               binding:"required"`
	Year        int        `json:"year"               gorm:"not null;default:null"                                               binding:"required,min=0"`
	Views       int        `json:"views"              gorm:"not null;default:0"`
	Cover       string     `json:"cover"              gorm:"default:null"`
	Authors     []*Author  `json:"authors,omitempty"  gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Tags        []*Tag     `json:"tags,omitempty"     gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Comments    []Comment  `json:"comments,omitempty" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	BookRates   []BookRate `json:"rates,omitempty"    gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Chapters    []Chapter  `json:"chapters,omitempty" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
