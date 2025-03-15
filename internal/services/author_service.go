package services

import (
	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type AuthorService struct {
	db *gorm.DB
}

func NewAuthorService(db *gorm.DB) AuthorService {
	return AuthorService{
		db: db,
	}
}

func (service AuthorService) GetAll() ([]models.Author, error) {
	var authors []models.Author
	tx := service.db.Find(&authors)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return authors, nil
}

func (service AuthorService) GetById(id uint) (models.Author, error) {
	var author models.Author
	tx := service.db.First(&author, id)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Create(author models.Author) (models.Author, error) {
	tx := service.db.Create(&author)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Update(id uint, newAuthor models.Author) (models.Author, error) {
	var author models.Author
	tx := service.db.First(&author, id)

	if tx.Error != nil {
		return author, tx.Error
	}

	author.Fullname = newAuthor.Fullname
	author.Biography = newAuthor.Biography
	author.Image = newAuthor.Image

	tx = service.db.Save(&author)

	if tx.Error != nil {
		return author, tx.Error
	}

	return author, nil
}

func (service AuthorService) Delete(id uint) error {
	var author models.Author
	tx := service.db.First(&author, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Delete(&author)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service AuthorService) GetBooks(id uint) ([]models.Book, error) {
	var author models.Author
	tx := service.db.First(&author, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var books []models.Book
	err := service.db.Model(&author).Association("Books").Find(&books)

	if err != nil {
		return nil, err
	}

	return books, nil

}

func (service AuthorService) AddBook(authorId uint, bookId uint) error {
	var author models.Author
	tx := service.db.First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	author.Books = append(author.Books, &book)
	tx = service.db.Save(&author)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service AuthorService) DeleteBook(authorId uint, bookId uint) error {
	var author models.Author
	tx := service.db.First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	err := service.db.Model(&author).Association("Books").Delete(&book)

	if err != nil {
		return err
	}

	return nil
}
