package models

import (
	"time"
)

type Chapter struct {
	Id        uint      `json:"id" gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title" gorm:"not null;default:null" binding:"required"`
	Content   string    `json:"content" gorm:"not null;default:null" binding:"required"`
	Serial    int       `json:"serial" gorm:"not null;default:null" binding:"required"`
	BookId    uint      `json:"book_id" gorm:"not null;default:null" binding:"required"`
}
