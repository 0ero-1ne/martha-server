package models

import (
	"time"
)

type Chapter struct {
	Id        uint      `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Serial    int       `json:"serial"`
	Date      time.Time `json:"date"`
	BookId    uint      `json:"book_id"`
}
