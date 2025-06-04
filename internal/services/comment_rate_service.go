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
	commentRate.User = models.User{}
	tx := service.db.Create(&commentRate)

	return commentRate, tx.Error
}

func (service CommentRateService) Update(newCommentRate models.CommentsRates) (models.CommentsRates, error) {
	var commentRate models.CommentsRates
	tx := service.db.
		Where("user_id = ? and comment_id = ?", newCommentRate.UserId, newCommentRate.CommentId).
		Find(&commentRate)

	if tx.Error != nil {
		return newCommentRate, tx.Error
	}

	tx = service.db.Model(&models.CommentsRates{}).
		Where("user_id = ? and comment_id = ?", commentRate.UserId, commentRate.CommentId).
		Update("rating", newCommentRate.Rating)

	return newCommentRate, tx.Error
}

func (service CommentRateService) Delete(commentId uint, userId uint) error {
	tx := service.db.Where("comment_id = ? and user_id = ?", commentId, userId).Delete(&models.CommentsRates{})

	if tx.RowsAffected == 0 {
		return errors.New("CommentRate not found")
	}

	return nil
}
