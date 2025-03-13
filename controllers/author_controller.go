package controllers

import (
	"net/http"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	service services.AuthorService
}

func NewAuthorController(service services.AuthorService) AuthorController {
	return AuthorController{
		service: service,
	}
}

func (controller AuthorController) GetAll(ctx *gin.Context) {
	authors, err := controller.service.GetAll()

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (controller AuthorController) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	author, err := controller.service.GetById(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (controller AuthorController) Create(ctx *gin.Context) {
	var author models.Author
	err := ctx.ShouldBindJSON(&author)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAuthor, err := controller.service.Create(author)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAuthor)
}

func (controller AuthorController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	var author models.Author
	err = ctx.ShouldBindJSON(&author)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAuthor, err := controller.service.Update(uint(id), author)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedAuthor)
}

func (controller AuthorController) Delete(ctx *gin.Context) {
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

func (controller AuthorController) GetBooks(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	books, err := controller.service.GetBooks(uint(bookId))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (controller AuthorController) AddBook(ctx *gin.Context) {
	authorId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	bookId, err := strconv.ParseUint(ctx.Param("book_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'book_id' param value"})
		return
	}

	err = controller.service.AddBook(uint(authorId), uint(bookId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller AuthorController) DeleteBook(ctx *gin.Context) {
	authorId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	bookId, err := strconv.ParseUint(ctx.Param("book_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'book_id' param value"})
		return
	}

	err = controller.service.DeleteBook(uint(authorId), uint(bookId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
