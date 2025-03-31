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

type RefreshResponse struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func NewAuthController(service services.AuthService, jwtManager utils.JWTManager) AuthController {
	return AuthController{
		service:    service,
		jwtManager: jwtManager,
	}
}

func (controller AuthController) Signup(ctx *gin.Context) {
	var authUser models.AuthUser

	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, "Incorrect email or password")
		return
	}

	if err := controller.service.Signup(authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "User was successully created")
}

func (controller AuthController) Login(ctx *gin.Context) {
	var authUser models.AuthUser

	if err := ctx.ShouldBindJSON(&authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, "Incorrect email or password")
		return
	}

	user, err := controller.service.Login(authUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Incorrect email or password")
		return
	}

	if !checkPasswordHash(authUser.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, "Incorrect email or password")
		return
	}

	accessToken, err := controller.jwtManager.NewJWTToken(user.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Login error, try again later")
		log.Printf("%s", err.Error())
		return
	}

	refreshToken, err := controller.jwtManager.NewRefreshToken(user.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Login error, try again later")
		log.Printf("%s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (controller AuthController) Refresh(ctx *gin.Context) {
	var refreshResponse RefreshResponse

	if err := ctx.ShouldBindJSON(&refreshResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, "No refresh token")
		return
	}

	refreshToken := refreshResponse.RefreshToken

	if err := controller.jwtManager.VerifyToken(refreshToken); !err {
		ctx.JSON(http.StatusUnauthorized, "Refresh token is expired")
		return
	}

	userId, err := controller.jwtManager.ExtractIdFromToken(refreshToken)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := controller.jwtManager.NewJWTToken(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Generate access token error")
		log.Printf("%s", err.Error())
		return
	}

	refreshToken, err = controller.jwtManager.NewRefreshToken(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Generate refresh token error")
		log.Printf("%s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
