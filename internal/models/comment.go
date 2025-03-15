package models

import "time"

type Comment struct {
	Id        uint          `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Text      string        `json:"text" gorm:"not null;default:null"`
	BookId    uint          `json:"book_id" gorm:"not null;default:null"`
	UserId    uint          `json:"user_id" gorm:"not null;default:null"`
	Rates     []CommentRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
}
