package models

import (
	"time"
)

type Tag struct {
	Id        uint      `json:"id" gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title" gorm:"uniqueIndex;not null;default:null" binding:"required,min=2"`
	Books     []*Book   `json:"books," gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
