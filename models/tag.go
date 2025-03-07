package models

import (
	"time"
)

type Tag struct {
	Id        uint      `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
	Title     string    `gorm:"uniqueIndex:not null;"`
	Books     []*Book   `gorm:"many2many:book_tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
