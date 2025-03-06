package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Fullname  string
	Biography string
	Books     []*Book `gorm:"many2many:book_authors;"`
}
