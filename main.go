package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prdpx7/go-crud/controllers"
	"github.com/prdpx7/go-crud/database"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", controllers.HealthCheck)
	r.GET("/ping", controllers.Ping)
	r.GET("/authors", controllers.ListAuthors)
	r.POST("/authors", controllers.CreateAuthor)
}

func main() {
	// follow readme to setup postgres and database first
	// Connect to local postgres with database name: go_crud
	database.SetupConnection()
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	setupRouter(app)
	app.Run(":8000")
}
