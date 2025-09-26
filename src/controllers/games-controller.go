package controllers

import (
	"chess-opening-analyzer/src/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGameByID(c *gin.Context) {
	id := c.Param("id")
	gameID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	gameRepository := repositories.NewGameRepository(db)
	game, err := gameRepository.GetGameByID(uint(gameID))

	if err != nil {
		c.JSON(404, gin.H{"error": "Game not found"})
	}

	c.JSON(200, game)
}
