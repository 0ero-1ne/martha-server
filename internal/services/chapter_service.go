package services

import (
	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/models"
)

type ChapterService struct {
	db *gorm.DB
}

func NewChapterService(db *gorm.DB) ChapterService {
	return ChapterService{
		db: db,
	}
}

func (service ChapterService) GetCount() int {
	var count int64
	tx := service.db.Model(&models.Chapter{}).Count(&count)
	if tx.Error != nil {
		return 0
	}
	return int(count)
}

func (service ChapterService) GetAll(params models.BookUrlParams) ([]models.Chapter, error) {
	var chapters []models.Chapter
	tx := service.db

	if params.Offset != 0 {
		tx = tx.Offset(params.Offset)
	}

	if params.Limit != 0 {
		tx = tx.Limit(params.Limit)
	}

	tx = tx.Order("book_id, serial").Find(&chapters)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return chapters, nil
}

func (service ChapterService) GetById(id uint) (models.Chapter, error) {
	var chapter models.Chapter
	tx := service.db.First(&chapter, id)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) GetChaptersByBookId(bookId uint) ([]models.Chapter, error) {
	var chapters []models.Chapter
	tx := service.db.Where("book_id = ?", bookId).Find(&chapters)

	return chapters, tx.Error
}

func (service ChapterService) Create(chapter models.Chapter) (models.Chapter, error) {
	var book models.Book
	tx := service.db.First(&book, chapter.BookId)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	tx = service.db.Create(&chapter)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) Update(id uint, newChapter models.Chapter) (models.Chapter, error) {
	var chapter models.Chapter
	tx := service.db.First(&chapter, id)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	var book models.Book
	tx = service.db.First(&book, chapter.BookId)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	chapter.Title = newChapter.Title
	chapter.Text = newChapter.Text
	chapter.Audio = newChapter.Audio
	chapter.BookId = newChapter.BookId
	chapter.Serial = newChapter.Serial

	tx = service.db.Save(&chapter)

	if tx.Error != nil {
		return chapter, tx.Error
	}

	return chapter, nil
}

func (service ChapterService) Delete(id uint) error {
	var chapter models.Chapter
	tx := service.db.First(&chapter, id)

	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Delete(&chapter)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
