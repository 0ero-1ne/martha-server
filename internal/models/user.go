package models

import "time"

type User struct {
	Id        uint       `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time  `                                                                        json:"created_at"`
	UpdatedAt time.Time  `                                                                        json:"updated_at"`
	Email     string     `gorm:"uniqueIndex;not null;default:null"                                json:"email"`
	Password  string     `gorm:"not null;default:null"                                            json:"password"`
	Username  string     `gorm:"not null;default:null"                                            json:"username"`
	Image     string     `                                                                        json:"image"`
	Comments  []Comment  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"                     json:"comments"`
	BookRates []BookRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"                     json:"rates"`
}
