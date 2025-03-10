package models

import (
	"time"
)

type Tag struct {
	Id        uint      `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `gorm:"uniqueIndex;not null;default:null" json:"title" binding:"required,min=2"`
	Books     []*Book   `gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"books,omitempty"`
}
