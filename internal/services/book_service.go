package services

import (
	"strings"

	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/utils"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return BookService{
		db: db,
	}
}

func (service BookService) GetAll(params models.BookUrlParams) ([]models.Book, error) {
	var books []models.Book
	tx := service.db.Model(&models.Book{})

	if params.WithAuthors {
		tx = tx.Preload("Authors")
	}

	if params.WithComments {
		tx = tx.Preload("Comments")
	}

	if params.WithTags {
		tx = tx.Preload("Tags")
	}

	if params.WithChapters {
		tx = tx.Preload("Chapters")
	}

	if params.WithBookRates {
		tx = tx.Preload("BooksRates")
	}

	if params.Offset != 0 {
		tx = tx.Offset(params.Offset)
	}

	if params.Limit != 0 {
		tx = tx.Limit(params.Limit)
	}

	if len(params.Query) != 0 {
		tx = tx.Where("lower(title) LIKE lower(?)", "%"+params.Query+"%")
	}

	if len(params.Tags) != 0 {
		tx = tx.Preload("Tags", "title in (?)", strings.Split(params.Tags, ","))
	}

	tx = tx.Find(&books)
	if len(params.Tags) != 0 {
		books = utils.Filter(books, func(book models.Book) bool {
			return len(book.Tags) != 0
		})
	}

	return books, tx.Error
}

func (service BookService) GetById(id uint, params models.BookUrlParams) (models.Book, error) {
	var book models.Book
	tx := service.db

	if params.WithAuthors {
		tx = tx.Preload("Authors")
	}

	if params.WithComments {
		tx = tx.Preload("Comments")
	}

	if params.WithTags {
		tx = tx.Preload("Tags")
	}

	if params.WithChapters {
		tx = tx.Preload("Chapters")
	}

	if params.WithBookRates {
		tx = tx.Preload("BooksRates")
	}

	tx = tx.First(&book, id)

	return book, tx.Error
}

func (service BookService) Create(book models.Book) (models.Book, error) {
	tx := service.db.Create(&book)

	if tx.Error != nil {
		return book, tx.Error
	}

	return book, nil
}

func (service BookService) Update(id uint, newBook models.Book) (models.Book, error) {
	var book models.Book
	tx := service.db.First(&book, id)

	if tx.Error != nil {
		return book, tx.Error
	}

	book.Title = newBook.Title
	book.Description = newBook.Description
	book.Status = newBook.Status
	book.Year = newBook.Year
	book.Views = newBook.Views
	book.Cover = newBook.Cover

	tx = service.db.Save(&book)

	if tx.Error != nil {
		return book, tx.Error
	}

	return book, nil
}

func (service BookService) Delete(id uint) error {
	var book models.Book
	tx := service.db.First(&book, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Delete(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// many2many book:tag

func (service BookService) GetTags(id uint) ([]models.Tag, error) {
	var book models.Book
	tx := service.db.First(&book, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var tags []models.Tag
	err := service.db.Model(&book).Association("Tags").Find(&tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (service BookService) AddTag(bookId uint, tagId uint) error {
	var book models.Book
	tx := service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var tag models.Tag
	tx = service.db.First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	book.Tags = append(book.Tags, &tag)
	tx = service.db.Save(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service BookService) DeleteTag(bookId uint, tagId uint) error {
	var book models.Book
	tx := service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var tag models.Tag
	tx = service.db.First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	err := service.db.Model(&book).Association("Tags").Delete(&tag)
	if err != nil {
		return err
	}

	return nil
}

// many2many book:author

func (service BookService) GetAuthors(id uint) ([]models.Author, error) {
	var book models.Book
	tx := service.db.First(&book, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var authors []models.Author
	err := service.db.Model(&book).Association("Authors").Find(&authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (service BookService) AddAuthor(bookId uint, authorId uint) error {
	var book models.Book
	tx := service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var author models.Author
	tx = service.db.First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	book.Authors = append(book.Authors, &author)
	tx = service.db.Save(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service BookService) DeleteAuthor(bookId uint, authorId uint) error {
	var book models.Book
	tx := service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var author models.Author
	tx = service.db.First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	err := service.db.Model(&book).Association("Authors").Delete(&author)
	if err != nil {
		return err
	}

	return nil
}
