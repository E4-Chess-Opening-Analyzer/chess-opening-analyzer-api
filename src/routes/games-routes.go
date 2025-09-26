package routes

import (
	"github.com/gin-gonic/gin"
	"chess-opening-analyzer/src/controllers"
)


// Get game by ID
// @Summary Get game by ID
// @Description Retrieve a game by its ID
// @Tags games
// @Accept json
// @Produce json
// @Param id path int true "Game ID"
// @Success 200 {object} models.Game
// @Failure 400
// @Failure 404
// @Router /games/{id} [get]
func GetGameByIdRoute(r *gin.RouterGroup) {
	r.GET("/:id", controllers.GetGameByID)
}

func SetupGameRoutes(r *gin.RouterGroup) {
	GetGameByIdRoute(r)
}