package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/prdpx7/go-crud/database"
)

type AuthorPayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

type BookPayload struct {
	AuthorID uint   `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

func CreateAuthor(ctx *gin.Context) {
	var payload AuthorPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author := database.Author{Name: payload.Name, Email: payload.Email}
	database.DB.Create(&author)
	ctx.JSON(http.StatusOK, gin.H{"data": author})
}

func ListAuthors(ctx *gin.Context) {
	var authors []database.Author
	database.DB.Find(&authors)
	ctx.JSON(http.StatusOK, gin.H{"data": authors})
}

func CreateBook(ctx *gin.Context) {
	var payload BookPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := database.Book{AuthorID: payload.AuthorID, Title: payload.Title}
	database.DB.Create(&book)
	database.DB.Joins("Author").Find(&book, book.ID)
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func ListBooks(ctx *gin.Context) {
	var books []database.Book
	database.DB.Joins("Author").Find(&books)
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
