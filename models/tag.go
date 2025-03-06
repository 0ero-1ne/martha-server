package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Title string
	Books []*Book `gorm:"many2many:book_tags;"`
}
