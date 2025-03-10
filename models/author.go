package models

import "time"

type Author struct {
	Id        uint      `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fullname  string    `json:"fullname" gorm:"not null;default:null"`
	Biography string    `json:"biography" gorm:";not null;default:null"`
	Image     string    `json:"image"`
	Books     []*Book   `gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"books"`
}
