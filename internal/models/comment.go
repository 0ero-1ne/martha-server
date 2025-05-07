package models

import "time"

type Comment struct {
	Id        uint          `gorm:"primarykey;not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"id"`
	CreatedAt time.Time     `                                                                        json:"created_at"`
	UpdatedAt time.Time     `                                                                        json:"updated_at"`
	Text      string        `gorm:"not null;default:null"                                            json:"text"`
	BookId    uint          `gorm:"not null;default:null"                                            json:"book_id"`
	UserId    uint          `gorm:"not null;default:null"                                            json:"user_id"`
	Rates     []CommentRate `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"                     json:"rates"`
	User      User          `json:"user"`
}
