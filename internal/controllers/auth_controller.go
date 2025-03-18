package controllers

import (
	"log"
	"net/http"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service    services.AuthService
	jwtManager utils.JWTManager
}

func NewAuthController(service services.AuthService, jwtManager utils.JWTManager) AuthController {
	return AuthController{
		service:    service,
		jwtManager: jwtManager,
	}
}

func (controller AuthController) SignUp(ctx *gin.Context) {
	var authUser models.AuthUser

	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.SignUp(authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (controller AuthController) SignIn(ctx *gin.Context) {
	var authUser models.AuthUser

	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	user, err := controller.service.SignIn(authUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	if !checkPasswordHash(authUser.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	accessToken, err := controller.jwtManager.NewJWTToken(user.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Login error, try again later"})
		log.Printf("%s", err.Error())
		return
	}

	refreshToken, err := controller.jwtManager.NewRefreshToken()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Login error, try again later"})
		log.Printf("%s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
