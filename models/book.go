package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title   string
	Authors []*Author `gorm:"many2many:book_authors;"`
	Tags    []*Tag    `gorm:"many2many:book_tags;"`
}
