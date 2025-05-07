package services

import (
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

func (service CommentService) GetAll() ([]models.Comment, error) {
	var comments []models.Comment
	tx := service.db.Find(&comments)

	return comments, tx.Error
}

func (service CommentService) GetById(commentId uint) (models.Comment, error) {
	var comment models.Comment
	tx := service.db.First(&comment, commentId)

	return comment, tx.Error
}

func (service CommentService) Create(comment models.Comment) (models.Comment, error) {
	tx := service.db.Create(&comment)
	return comment, tx.Error
}

func (service CommentService) Update(commentId uint, newComment models.Comment) (models.Comment, error) {
	var comment models.Comment

	tx := service.db.First(&comment, commentId)
	if tx.Error != nil {
		return comment, tx.Error
	}

	comment.Text = newComment.Text
	comment.BookId = newComment.BookId
	comment.UserId = newComment.UserId

	tx = service.db.Save(&comment)

	return comment, tx.Error
}

func (service CommentService) Delete(commentId uint) error {
	var comment models.Comment
	tx := service.db.First(&comment, commentId)

	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Delete(&comment)

	return tx.Error
}

// book

func (service CommentService) GetAllByBookId(bookId uint) ([]models.Comment, error) {
	var comments []models.Comment
	tx := service.db.
		Preload("Rates").
		Preload("Rates.User").
		Preload("User").
		Where("book_id = ?", bookId).
		Find(&comments)

	return comments, tx.Error
}
