package models

import "time"

type Comment struct {
	Id        uint          `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Text      string        `json:"text"`
	BookId    uint          `json:"book_id"`
	UserId    uint          `json:"user_id"`
	Rates     []CommentRate `json:"rates"`
	User      User          `json:"user"`
}
