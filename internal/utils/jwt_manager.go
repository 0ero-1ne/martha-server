package utils

import (
	"fmt"
	"strconv"
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

func (manager JWTManager) NewRefreshToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	sign, err := token.SignedString([]byte(manager.secret))

	if err != nil {
		return "", err
	}

	return sign, nil
}

func (manager JWTManager) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(manager.secret), nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

func (manager JWTManager) ExtractIdFromToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(manager.secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, fmt.Errorf("Invalid Token")
	}

	claimsSubToString := strconv.FormatFloat(claims["sub"].(float64), 'f', 0, 64)
	id, _ := strconv.ParseUint(claimsSubToString, 10, 64)

	return uint(id), nil
}
