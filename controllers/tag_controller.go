package controllers

import (
	"net/http"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	service services.TagService
}

func NewTagController(service services.TagService) TagController {
	return TagController{
		service: service,
	}
}

func (controller TagController) GetAll(ctx *gin.Context) {
	tags, err := controller.service.GetAll()

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (controller TagController) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	tag, err := controller.service.GetById(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (controller TagController) Create(ctx *gin.Context) {
	var tag models.Tag
	err := ctx.ShouldBindJSON(&tag)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTag, err := controller.service.Create(tag)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdTag)
}

func (controller TagController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	var tag models.Tag
	err = ctx.ShouldBindJSON(&tag)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTag, err := controller.service.Update(uint(id), tag)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedTag)
}

func (controller TagController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	err = controller.service.Delete(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}
