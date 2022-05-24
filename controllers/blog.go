package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/prdpx7/go-crud/database"
)

type AuthorPayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
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
