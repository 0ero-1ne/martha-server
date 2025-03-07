package models

import "time"

type User struct {
	Id        uint       `gorm:"primarykey;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"upated_at"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Username  string     `json:"username"`
	Image     string     `json:"image"`
	Comments  []Comment  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments"`
	BookRates []BookRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"rates"`
}
