package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/0ero-1ne/martha-server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return AuthService{
		db: db,
	}
}

func (service AuthService) Signup(authUser models.AuthUser) error {
	tx := service.db.First(&models.User{}, "email = ?", authUser.Email)

	if tx.Error == nil {
		return fmt.Errorf("Email %s already in use", authUser.Email)
	}

	hashedPassword, err := hashPassword(authUser.Password)

	if err != nil {
		return errors.New("Sign up error. Try again later")
	}

	user := models.User{
		Email:    strings.ToLower(authUser.Email),
		Username: authUser.Email,
		Password: hashedPassword,
	}

	tx = service.db.Save(&user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service AuthService) Login(authUser models.AuthUser) (models.User, error) {
	var user models.User
	tx := service.db.First(&user, "email = ?", authUser.Email)

	return user, tx.Error
}

func (service AuthService) Refresh() {

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
