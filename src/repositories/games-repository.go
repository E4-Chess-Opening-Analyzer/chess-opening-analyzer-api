package repositories

import (
	"chess-opening-analyzer/src/models"
	"fmt"

	"gorm.io/gorm"
)

type GamesRepository struct {
	DB *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GamesRepository {
	return &GamesRepository{DB: db}
}

func (repo *GamesRepository) GetGameByID(id uint) (*models.Game, error) {
	var game models.Game
    fmt.Println("Querying game with ID:", id)
	if err := repo.DB.First(&game, id).Error; err != nil {
		fmt.Println("Error fetching game:", err)
		return nil, err
	}
	return &game, nil
}