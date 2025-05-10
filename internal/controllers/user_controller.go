package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
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

func (controller UserController) Update(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.GetUint("user_id")
	if userId != user.Id {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	updatedUser, err := controller.service.Update(user, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
