package main

import (
	"chess-opening-analyzer/src/database"
	"chess-opening-analyzer/src/docs"
	"chess-opening-analyzer/src/middlewares"
	"log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	v1 := r.Group("/api")

	r.Use(middlewares.DatabaseMiddleware(database.DB))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
