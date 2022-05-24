package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "go-crud is up and running"})
}
