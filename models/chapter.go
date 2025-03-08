package models

import (
	"time"
)

type Chapter struct {
	Id        uint      `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
	Title     string    `json:"title" gorm:"not null;default:null"`
	Content   string    `json:"content" gorm:"not null;default:null"`
	Serial    int       `json:"serial" gorm:"not null;default:null"`
	Date      time.Time `json:"date" gorm:"not null;default:null"`
	BookId    uint      `json:"book_id" gorm:"not null;default:null"`
}
