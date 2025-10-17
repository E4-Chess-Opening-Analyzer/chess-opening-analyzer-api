package controllers

import (
	"chess-opening-analyzer/src/middlewares"
	"chess-opening-analyzer/src/repositories"
	"github.com/gin-gonic/gin"
)

func GetOutcomesForMoves(c *gin.Context) {
	var moves []string
	if err := c.ShouldBindJSON(&moves); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	db := middlewares.GetMongoDB(c)
	if db == nil {
		c.JSON(500, gin.H{"error": "Database not found"})
		return
	}

	gameRepo := repositories.NewMovesRepository(db)
	outcomes, err := gameRepo.GetOutcomesForMoves(moves)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get outcomes"})
		return
	}

	c.JSON(200, outcomes)
}
