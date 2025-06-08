package services

import (
	"errors"

	"github.com/0ero-1ne/martha-server/internal/models"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) CommentService {
	return CommentService{
		db: db,
	}
}

func (service CommentService) GetAll(userId uint) ([]models.Comment, error) {
	var comments []models.Comment
	tx := service.db
	if userId != 0 {
		tx = tx.Where("user_id = ?", userId).Preload("Rates")
	}
	tx = tx.Find(&comments)

	return comments, tx.Error
}

func (service CommentService) GetById(commentId uint) (models.Comment, error) {
	var comment models.Comment
	tx := service.db.First(&comment, commentId)

	return comment, tx.Error
}

func (service CommentService) Create(comment models.Comment) (models.Comment, error) {
	if *comment.ParentId == 0 {
		comment.ParentId = nil
	}
	tx := service.db.Create(&comment)
	return comment, tx.Error
}

func (service CommentService) Update(commentId uint, newComment models.Comment, userId uint) (models.Comment, error) {
	var comment models.Comment
	tx := service.db.Preload("Rates").Preload("User").First(&comment, commentId)
	if tx.Error != nil {
		return comment, tx.Error
	}

	var user models.User
	tx = service.db.First(&user, userId)
	if tx.Error != nil {
		return comment, errors.New("you have no right to update this comment")
	}

	if comment.UserId != userId && (user.Role != "moderator" && user.Role != "admin") {
		return comment, errors.New("you have no right to update this comment")
	}

	tx = service.db.Model(&models.Comment{}).
		Where("id = ?", comment.Id).
		Update("text", newComment.Text)

	return newComment, tx.Error
}

func (service CommentService) Delete(commentId uint, userId uint) error {
	var comment models.Comment
	tx := service.db.First(&comment, commentId)

	if tx.Error != nil {
		return tx.Error
	}

	var user models.User
	tx = service.db.First(&user, userId)
	if tx.Error != nil {
		return errors.New("you have no right to delete this comment")
	}

	if comment.UserId != userId && (user.Role != "moderator" && user.Role != "admin") {
		return errors.New("you have no right to delete this comment")
	}

	tx = service.db.Delete(&comment)

	return tx.Error
}

// book

func (service CommentService) GetAllByBookId(bookId uint) ([]models.Comment, error) {
	var comments []models.Comment
	tx := service.db.
		Model(&models.Comment{}).
		Where("book_id = ? and parent_id is null", bookId).
		Preload("User").
		Preload("Rates").
		Preload("Rates.User").
		Preload("Replies", preloadReplies).
		Find(&comments)

	return comments, tx.Error
}

func preloadReplies(d *gorm.DB) *gorm.DB {
	return d.Preload("User").Preload("Rates.User").Preload("Replies", preloadReplies)
}
