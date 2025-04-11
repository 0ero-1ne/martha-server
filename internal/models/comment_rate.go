package models

type CommentRate struct {
	CommentId uint `gorm:"primarykey;not null;"   json:"comment_id"`
	UserId    uint `gorm:"primarykey;not null;"   json:"user_id"`
	Rating    bool `gorm:"not null;default:null;" json:"rating"` // true = +1, false = -1
}
