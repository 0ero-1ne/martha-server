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

func (service CommentRateService) GetAll() ([]models.CommentsRates, error) {
	var commentRates []models.CommentsRates
	tx := service.db.Find(&commentRates)

	return commentRates, tx.Error
}

func (service CommentRateService) Create(commentRate models.CommentsRates) (models.CommentsRates, error) {
	tx := service.db.Save(&commentRate)

	return commentRate, tx.Error
}

func (service CommentRateService) Update(newCommentRate models.CommentsRates) (models.CommentsRates, error) {
	var commentRate models.CommentsRates
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

func (service CommentRateService) Delete(commentRate models.CommentsRates) error {
	tx := service.db.Delete(&commentRate)
	if tx.RowsAffected == 0 {
		return errors.New("CommentRate not found")
	}

	return nil
}
