package main

import (
	"chess-opening-analyzer/src/database"
	"chess-opening-analyzer/src/docs"
	"chess-opening-analyzer/src/middlewares"
	"chess-opening-analyzer/src/models"
	"chess-opening-analyzer/src/routes"
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

	// Auto-migrate the database schema
    if err := database.DB.AutoMigrate(&models.Game{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	// Apply database middleware
	r.Use(middlewares.DatabaseMiddleware(database.DB))

	v1 := r.Group("/api")
	routes.SetupGameRoutes(v1.Group("/games"))

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
