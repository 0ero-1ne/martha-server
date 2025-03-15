package models

import "time"

type User struct {
	Id        uint       `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null;default:null"`
	Password  string     `json:"password" gorm:"not null;default:null"`
	Username  string     `json:"username" gorm:"not null;default:null"`
	Image     string     `json:"image"`
	Comments  []Comment  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments"`
	BookRates []BookRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
}
