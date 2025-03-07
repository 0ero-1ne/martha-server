package models

type BookRate struct {
	BookId uint `gorm:"primarykey;" json:"book_id"`
	UserId uint `gorm:"primarykey;" json:"user_id"`
	Rating int8 `gorm:"check:rating between 1 and 5;" json:"rating"`
}
