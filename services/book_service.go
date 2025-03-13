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

// many2many book:tag

func (service BookService) GetTags(id uint) ([]models.Tag, error) {
	var book models.Book
	tx := db.GetDB().First(&book, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var tags []models.Tag
	err := db.GetDB().Model(&book).Association("Tags").Find(&tags)

	if err != nil {
		return nil, err
	}

	return tags, nil

}

func (service BookService) AddTag(bookId uint, tagId uint) error {
	var book models.Book
	tx := db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var tag models.Tag
	tx = db.GetDB().First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	book.Tags = append(book.Tags, &tag)
	tx = db.GetDB().Save(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service BookService) DeleteTag(bookId uint, tagId uint) error {
	var book models.Book
	tx := db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var tag models.Tag
	tx = db.GetDB().First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	err := db.GetDB().Model(&book).Association("Tags").Delete(&tag)

	if err != nil {
		return err
	}

	return nil
}

// many2many book:author

func (service BookService) GetAuthors(id uint) ([]models.Author, error) {
	var book models.Book
	tx := db.GetDB().First(&book, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var authors []models.Author
	err := db.GetDB().Model(&book).Association("Authors").Find(&authors)

	if err != nil {
		return nil, err
	}

	return authors, nil

}

func (service BookService) AddAuthor(bookId uint, authorId uint) error {
	var book models.Book
	tx := db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var author models.Author
	tx = db.GetDB().First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	book.Authors = append(book.Authors, &author)
	tx = db.GetDB().Save(&book)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service BookService) DeleteAuthor(bookId uint, authorId uint) error {
	var book models.Book
	tx := db.GetDB().First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	var author models.Author
	tx = db.GetDB().First(&author, authorId)

	if tx.Error != nil {
		return tx.Error
	}

	err := db.GetDB().Model(&book).Association("Authors").Delete(&author)

	if err != nil {
		return err
	}

	return nil
}
