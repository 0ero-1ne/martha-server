package services

import (
	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type TagService struct {
	db *gorm.DB
}

func NewTagService(db *gorm.DB) TagService {
	return TagService{
		db: db,
	}
}

func (service TagService) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	tx := service.db.Find(&tags)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

func (service TagService) GetById(id uint) (models.Tag, error) {
	var tag models.Tag
	tx := service.db.First(&tag, id)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Create(tag models.Tag) (models.Tag, error) {
	tx := service.db.Create(&tag)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Update(id uint, newTag models.Tag) (models.Tag, error) {
	var tag models.Tag
	tx := service.db.First(&tag, id)

	if tx.Error != nil {
		return tag, tx.Error
	}

	tag.Title = newTag.Title

	tx = service.db.Save(&tag)

	if tx.Error != nil {
		return tag, tx.Error
	}

	return tag, nil
}

func (service TagService) Delete(id uint) error {
	var tag models.Tag
	tx := service.db.First(&tag, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Delete(&tag)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service TagService) GetBooks(id uint) ([]models.Book, error) {
	var tag models.Tag
	tx := service.db.First(&tag, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var books []models.Book
	err := service.db.Model(&tag).Association("Books").Find(&books)

	if err != nil {
		return nil, err
	}

	return books, nil

}

func (service TagService) AddBook(tagId uint, bookId uint) error {
	var tag models.Tag
	tx := service.db.First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	tag.Books = append(tag.Books, &book)
	tx = service.db.Save(&tag)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service TagService) DeleteBook(tagId uint, bookId uint) error {
	var tag models.Tag
	tx := service.db.First(&tag, tagId)

	if tx.Error != nil {
		return tx.Error
	}

	var book models.Book
	tx = service.db.First(&book, bookId)

	if tx.Error != nil {
		return tx.Error
	}

	err := service.db.Model(&tag).Association("Books").Delete(&book)

	if err != nil {
		return err
	}

	return nil
}
