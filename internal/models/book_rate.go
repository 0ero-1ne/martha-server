package models

type BookRate struct {
	BookId uint `gorm:"primarykey;not null;"                   json:"book_id"`
	UserId uint `gorm:"primarykey;not null;"                   json:"user_id"`
	Rating int8 `gorm:"check:rating between 1 and 5;not null;" json:"rating"`
}
