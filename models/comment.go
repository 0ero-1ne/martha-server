package models

import "time"

type Comment struct {
	Id        uint          `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"upated_at"`
	Text      string        `json:"text"`
	BookId    uint          `json:"book_id"`
	UserId    uint          `json:"user_id"`
	Rates     []CommentRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
}
