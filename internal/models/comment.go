package models

import "time"

type Comment struct {
	Id        uint            `json:"id"`
	ParentId  *uint           `json:"parent_id,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Text      string          `json:"text"`
	BookId    uint            `json:"book_id"`
	UserId    uint            `json:"user_id"`
	Rates     []CommentsRates `json:"rates,omitempty"`
	Replies   []Comment       `json:"replies,omitempty" gorm:"foreignKey:ParentId"`
	User      User            `json:"user" gorm:"foreignKey:UserId"`
}
