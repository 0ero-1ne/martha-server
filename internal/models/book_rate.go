package models

type BookRate struct {
	BookId uint `json:"book_id"`
	UserId uint `json:"user_id"`
	Rating int8 `json:"rating"`
}
