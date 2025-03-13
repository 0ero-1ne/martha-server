package services

import (
	"server/db"
	"server/models"
)

type ChapterService struct{}

func (service ChapterService) GetAll() ([]models.Chapter, error) {
	var chapters []models.Chapter
	tx := db.GetDB().Find(&chapters)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return chapters, nil
}

func (service ChapterService) GetById(id uint) (models.Chapter, error) {
	var chapter models.Chapter
	tx := db.GetDB().First(&chapter, id)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) Create(chapter models.Chapter) (models.Chapter, error) {
	var book models.Book
	tx := db.GetDB().First(&book, chapter.BookId)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	tx = db.GetDB().Create(&chapter)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) Update(id uint, newChapter models.Chapter) (models.Chapter, error) {
	var chapter models.Chapter
	tx := db.GetDB().First(&chapter, id)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	var book models.Book
	tx = db.GetDB().First(&book, chapter.BookId)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	chapter.Title = newChapter.Title
	chapter.Content = newChapter.Content
	chapter.BookId = newChapter.BookId
	chapter.Serial = newChapter.Serial

	tx = db.GetDB().Save(&chapter)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) Delete(id uint) error {
	var chapter models.Chapter
	tx := db.GetDB().First(&chapter, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = db.GetDB().Delete(&chapter)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
