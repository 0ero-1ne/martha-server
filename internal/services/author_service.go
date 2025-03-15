package services

import (
	"github.com/0ero-1ne/martha/internal/db"
	"github.com/0ero-1ne/martha/internal/models"
)

type AuthorService struct{}

func (service AuthorService) GetAll() ([]models.Author, error) {
	var authors []models.Author
	tx := db.GetDB().Find(&authors)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return authors, nil
}

func (service AuthorService) GetById(id uint) (models.Author, error) {
	var author models.Author
	tx := db.GetDB().First(&author, id)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Create(author models.Author) (models.Author, error) {
	tx := db.GetDB().Create(&author)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Update(id uint, newAuthor models.Author) (models.Author, error) {
	var author models.Author
	tx := db.GetDB().First(&author, id)

	if tx.Error != nil {
		return author, tx.Error
	}

	author.Fullname = newAuthor.Fullname
	author.Biography = newAuthor.Biography
	author.Image = newAuthor.Image

	tx = db.GetDB().Save(&author)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Delete(id uint) error {
	var author models.Author
	tx := db.GetDB().First(&author, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = db.GetDB().Delete(&author)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service AuthorService) GetBooks(id uint) ([]models.Book, error) {
	var author models.Author
	tx := db.GetDB().First(&author, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var books []models.Book
	err := db.GetDB().Model(&author).Association("Books").Find(&books)

	if err != nil {
		return nil, err
	}

	return books, nil

}

func (service AuthorService) AddBook(authorId uint, bookId uint) error {
	var author models.Author
	tx := db.GetDB().First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	author.Books = append(author.Books, &book)
	tx = db.GetDB().Save(&author)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service AuthorService) DeleteBook(authorId uint, bookId uint) error {
	var author models.Author
	tx := db.GetDB().First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	err := db.GetDB().Model(&author).Association("Books").Delete(&book)

	if err != nil {
		return err
	}

	return nil
}
