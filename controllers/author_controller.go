package controllers

import (
	"net/http"
	"server/models"
	"server/services"

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
	authorId := ctx.GetUint("author_id")
	author, err := controller.service.GetById(authorId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (controller AuthorController) Create(ctx *gin.Context) {
	var author models.Author

	if err := ctx.ShouldBindJSON(&author); err != nil {
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
	authorId := ctx.GetUint("author_id")
	var author models.Author

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAuthor, err := controller.service.Update(authorId, author)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedAuthor)
}

func (controller AuthorController) Delete(ctx *gin.Context) {
	authorId := ctx.GetUint("author_id")

	if err := controller.service.Delete(authorId); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many author:book

func (controller AuthorController) GetBooks(ctx *gin.Context) {
	authorId := ctx.GetUint("author_id")
	books, err := controller.service.GetBooks(authorId)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (controller AuthorController) AddBook(ctx *gin.Context) {
	authorId := ctx.GetUint("author_id")
	bookId := ctx.GetUint("book_id")

	if err := controller.service.AddBook(authorId, bookId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller AuthorController) DeleteBook(ctx *gin.Context) {
	authorId := ctx.GetUint("author_id")
	bookId := ctx.GetUint("book_id")

	if err := controller.service.DeleteBook(authorId, bookId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
