package services

import (
	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{
		db: db,
	}
}

func (service UserService) GetById(id uint) (models.User, error) {
	var user models.User
	tx := service.db.First(&user, id)

	return user, tx.Error
}
