package services

import (
	"server/db"
	"server/models"
)

type BookService struct{}

func NewBookService() BookService {
	return BookService{}
}

func (service BookService) GetAll() ([]models.Book, error) {
	var books []models.Book
	tx := db.GetDB().Find(&books)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return books, nil
}

func (service BookService) GetById(id uint) (models.Book, error) {
	var book models.Book
	tx := db.GetDB().First(&book, id)

	if tx.Error != nil {
		return book, tx.Error
	}

	return book, nil
}

func (service BookService) Create(book models.Book) (models.Book, error) {
	tx := db.GetDB().Create(&book)

	if tx.Error != nil {
		return book, tx.Error
	}

	return book, nil
}

func (service BookService) Update(id uint, newBook models.Book) (models.Book, error) {
	var book models.Book
	tx := db.GetDB().First(&book, id)

	if tx.Error != nil {
		return book, tx.Error
	}

	book.Title = newBook.Title
	book.Description = newBook.Description
	book.Status = newBook.Status
	book.Year = newBook.Year
	book.Views = newBook.Views
	book.Cover = newBook.Cover

	tx = db.GetDB().Save(&book)

	if tx.Error != nil {
		return book, tx.Error
	}

	return book, nil
}

func (service BookService) Delete(id uint) error {
	var book models.Book
	tx := db.GetDB().First(&book, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = db.GetDB().Delete(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
