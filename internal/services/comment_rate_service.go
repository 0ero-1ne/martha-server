package services

import (
	"errors"

	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type CommentRateService struct {
	db *gorm.DB
}

func NewCommentRateService(db *gorm.DB) CommentRateService {
	return CommentRateService{
		db: db,
	}
}

func (service CommentRateService) GetAll() ([]models.CommentRate, error) {
	var commentRates []models.CommentRate
	tx := service.db.Find(&commentRates)

	return commentRates, tx.Error
}

func (service CommentRateService) Create(commentRate models.CommentRate) (models.CommentRate, error) {
	tx := service.db.Save(&commentRate)

	return commentRate, tx.Error
}

func (service CommentRateService) Update(newCommentRate models.CommentRate) (models.CommentRate, error) {
	var commentRate models.CommentRate
	tx := service.db.
		Where("user_id = ?", newCommentRate.UserId).
		Where("comment_id = ?", newCommentRate.CommentId).
		Find(&commentRate)

	if tx.Error != nil {
		return newCommentRate, tx.Error
	}

	commentRate.Rating = newCommentRate.Rating

	tx = service.db.Save(&commentRate)

	return commentRate, tx.Error
}

func (service CommentRateService) Delete(commentRate models.CommentRate) error {
	tx := service.db.Delete(&commentRate)
	if tx.RowsAffected == 0 {
		return errors.New("CommentRate not found")
	}

	return nil
}
