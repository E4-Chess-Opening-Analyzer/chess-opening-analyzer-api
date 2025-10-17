package main

import (
	"chess-opening-analyzer/src/database"
	"chess-opening-analyzer/src/docs"
	"chess-opening-analyzer/src/middlewares"
	"chess-opening-analyzer/src/routes"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Connect to MongoDB
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err := db.Client().Ping(context.Background(), nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	// Apply MongoDB middleware
	r.Use(middlewares.MongoMiddleware(db))

	v1 := r.Group("/api")
	routes.SetupGameRoutes(v1.Group("/games"))

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
