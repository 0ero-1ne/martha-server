package models

type CommentRate struct {
	CommentId uint `gorm:"primarykey;" json:"comment_id"`
	UserId    uint `gorm:"primarykey;" json:"user_id"`
	Rating    bool `json:"rating"` // true = +1, false = -1
}
