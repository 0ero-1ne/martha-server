package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/0ero-1ne/martha-server/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret string
}

func NewJWTManager(config config.JWTConfig) JWTManager {
	return JWTManager{
		secret: config.Secret,
	}
}

func (manager JWTManager) NewJWTToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	sign, err := token.SignedString([]byte(manager.secret))

	if err != nil {
		return "", err
	}

	return sign, nil
}

func (manager JWTManager) NewRefreshToken() (string, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (manager JWTManager) Parse(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Token is invalid")
	}

	return nil
}
