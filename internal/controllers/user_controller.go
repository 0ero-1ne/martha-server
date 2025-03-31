package controllers

import (
	"net/http"

	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserController{
		service: service,
	}
}

func (controller UserController) GetById(ctx *gin.Context) {
	userId := ctx.GetUint("user_id")
	user, err := controller.service.GetById(userId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
