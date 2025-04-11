package models

import "time"

type Author struct {
	Id        uint      `json:"id"              gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fullname  string    `json:"fullname"        gorm:"not null;default:null"                                               binding:"required"`
	Biography string    `json:"biography"       gorm:";not null;default:null"                                              binding:"required"`
	Image     string    `json:"image"`
	Books     []*Book   `json:"books,omitempty" gorm:"many2many:book_authors;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
