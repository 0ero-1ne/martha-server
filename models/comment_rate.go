package models

type CommentRate struct {
	CommentId uint `gorm:"primarykey;not null;" json:"comment_id"`
	UserId    uint `gorm:"primarykey;not null;" json:"user_id"`
	Rating    bool `json:"rating" gorm:"not null;default:null;"` // true = +1, false = -1
}
