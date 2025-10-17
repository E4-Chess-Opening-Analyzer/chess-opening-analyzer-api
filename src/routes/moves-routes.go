package routes

import (
	"github.com/gin-gonic/gin"
	"chess-opening-analyzer/src/controllers"
)


// Get outcomes for moves
// @Summary Get outcomes for moves
// @Description Retrieve outcomes using the moves played
// @Tags games
// @Accept json
// @Produce json
// @Param moves body models.Moves true "Moves object"
// @Success 201 {object} models.Outcome
// @Failure 400
// @Failure 500
// @Router /games/outcomes [post]
func GetOutcomesForMoves(r *gin.RouterGroup) {
	r.POST("", controllers.GetOutcomesForMoves)
}


func SetupGameRoutes(r *gin.RouterGroup) {
	GetOutcomesForMoves(r)
}