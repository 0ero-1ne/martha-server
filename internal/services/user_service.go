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

func (service UserService) GetCount() int {
	var count int64
	tx := service.db.Model(&models.User{}).Count(&count)
	if tx.Error != nil {
		return 0
	}
	return int(count)
}

func (service UserService) GetAll(params models.BookUrlParams) ([]models.User, error) {
	var users []models.User
	tx := service.db

	if params.Offset != 0 {
		tx = tx.Offset(params.Offset)
	}

	if params.Limit != 0 {
		tx = tx.Limit(params.Limit)
	}

	tx = tx.Order("id asc").Find(&users)

	return users, tx.Error
}

func (service UserService) MakeModer(id uint) error {
	var user models.User
	tx := service.db.First(&user, id)

	if tx.Error != nil {
		return tx.Error
	}

	user.Role = "moderator"

	tx = service.db.Save(&user)
	return tx.Error
}

func (service UserService) MakeUser(id uint) error {
	var user models.User
	tx := service.db.First(&user, id)

	if tx.Error != nil {
		return tx.Error
	}

	user.Role = "user"

	tx = service.db.Save(&user)
	return tx.Error
}

func (service UserService) GetById(id uint) (models.User, error) {
	var user models.User
	tx := service.db.First(&user, id)

	return user, tx.Error
}

func (service UserService) Update(newUser models.User, userId uint) (models.User, error) {
	var user models.User
	tx := service.db.First(&user, userId)

	if tx.Error != nil {
		return user, tx.Error
	}

	user.Email = newUser.Email
	user.Username = newUser.Username
	user.Image = newUser.Image
	user.Role = newUser.Role
	user.SavedBooks = newUser.SavedBooks

	tx = service.db.Save(&user)
	return user, tx.Error
}
