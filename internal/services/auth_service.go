package services

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/models"
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
		return fmt.Errorf("email %s already in use", authUser.Email)
	}

	hashedPassword, err := hashPassword(authUser.Password)
	if err != nil {
		return errors.New("sign up error. Try again later")
	}

	user := models.User{
		Email:    strings.ToLower(authUser.Email),
		Username: authUser.Email,
		Password: hashedPassword,
		Role:     "user",
		SavedBooks: models.SavedBooks{
			models.Reading:   []models.SavedBook{},
			models.Ended:     []models.SavedBook{},
			models.Stopped:   []models.SavedBook{},
			models.Planed:    []models.SavedBook{},
			models.Favorites: []models.SavedBook{},
		},
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

func (service AuthService) ChangePassword(passwords models.AuthPasswords, userId uint) error {
	var user models.User
	tx := service.db.First(&user, userId)
	if tx.Error != nil {
		return errors.New("Wrong old password")
	}

	if err := checkPasswordHash(passwords.OldPassword, user.Password); err == false {
		return errors.New("Wrong old password")
	}

	hashedNewPassword, err := hashPassword(passwords.NewPassword)
	if err != nil {
		return errors.New("Server error, try again later")
	}

	user.Password = hashedNewPassword
	user.Email = user.Email // костыль для Save

	tx = service.db.Save(&user)
	if tx.Error != nil {
		return errors.New("Server error, try again later")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
