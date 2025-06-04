package models

type CommentsRates struct {
	CommentId uint `json:"comment_id"`
	UserId    uint `json:"user_id"`
	Rating    bool `json:"rating"` // true = +1, false = -1
	User      User `json:"user"`
}
