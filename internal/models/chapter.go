package models

import (
	"time"
)

type Chapter struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title" binding:"required"`
	Text      string    `json:"text"`
	Audio     string    `json:"audio"`
	Serial    int       `json:"serial" binding:"required"`
	BookId    uint      `json:"book_id" binding:"required"`
}
