package repositories

import (
	"chess-opening-analyzer/src/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovesRepository struct {
	DB *mongo.Database
}

func NewMovesRepository(db *mongo.Database) *MovesRepository {
	return &MovesRepository{DB: db}
}


// Outcome = map[nextMove]map[result]int
func (repo *MovesRepository) GetOutcomesForMoves(moves []string) (models.Outcome, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rootNode map[string]*models.Moves
	err := repo.DB.Collection("opening").FindOne(ctx, bson.M{}).Decode(&rootNode)
	if err != nil {
		return nil, err
	}

	current := rootNode
	for _, move := range moves {
		node, exists := current[move]
		if !exists {
			return nil, errors.New("moves not found in database")
		}
		current = node.Next
	}

	outcomes := make(models.Outcome)
	for nextMove, node := range current {
		outcomes[nextMove] = map[int]uint{
			1:  uint(node.WhiteWin),
			0:  uint(node.Draw),
			-1: uint(node.BlackWin),
		}
	}

	return outcomes, nil
}