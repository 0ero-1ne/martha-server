package controllers

import (
	"net/http"
	"strconv"

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

func (controller UserController) GetCount(ctx *gin.Context) {
	count := controller.service.GetCount()
	ctx.AbortWithStatusJSON(http.StatusOK, count)
}

func (controller UserController) GetAll(ctx *gin.Context) {
	params := models.BookUrlParams{}

	if offset, offsetErr := strconv.ParseInt(ctx.Query("offset"), 10, 32); offsetErr == nil {
		params.Offset = int(offset)
	}

	if limit, limitErr := strconv.ParseInt(ctx.Query("limit"), 10, 32); limitErr == nil {
		params.Limit = int(limit)
	}

	users, err := controller.service.GetAll(params)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (controller UserController) MakeModer(ctx *gin.Context) {
	userId := ctx.GetUint("user_id")
	if err := controller.service.MakeModer(uint(userId)); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}

func (controller UserController) MakeUser(ctx *gin.Context) {
	userId := ctx.GetUint("user_id")
	if err := controller.service.MakeUser(uint(userId)); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
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
