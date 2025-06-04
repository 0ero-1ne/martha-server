package services

import (
	"errors"

	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type BookRateService struct {
	db *gorm.DB
}

func NewBookRateService(db *gorm.DB) BookRateService {
	return BookRateService{
		db: db,
	}
}

func (service BookRateService) GetAll() ([]models.BooksRates, error) {
	var bookRates []models.BooksRates
	tx := service.db.Find(&bookRates)

	return bookRates, tx.Error
}

func (service BookRateService) Create(bookRate models.BooksRates) (models.BooksRates, error) {
	bookRate.User = models.User{}
	tx := service.db.Create(&bookRate)

	return bookRate, tx.Error
}

func (service BookRateService) Update(newBookRate models.BooksRates) (models.BooksRates, error) {
	var bookRate models.BooksRates
	tx := service.db.
		Where("user_id = ? and book_id = ?", newBookRate.UserId, newBookRate.BookId).
		Find(&bookRate)

	if tx.Error != nil {
		return newBookRate, tx.Error
	}

	tx = service.db.Model(&models.BooksRates{}).
		Where("user_id = ? and book_id = ?", bookRate.UserId, bookRate.BookId).
		Update("rating", newBookRate.Rating)

	return newBookRate, tx.Error
}

func (service BookRateService) Delete(bookId uint, userId uint) error {
	tx := service.db.Where("book_id = ? and user_id = ?", bookId, userId).Delete(&models.BooksRates{})

	if tx.RowsAffected == 0 {
		return errors.New("BookRate not found")
	}

	return nil
}
