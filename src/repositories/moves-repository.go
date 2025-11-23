package repositories

import (
	"chess-opening-analyzer/src/models"
	"context"
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

// GetOutcomesForMoves Outcome = map[nextMove]map[result]int
func (repo *MovesRepository) GetOutcomesForMoves(moves []string) (models.Outcome, error) {
	collection := repo.DB.Collection("openings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// On récupère le document correspondant à la séquence
	filter := bson.M{"move_sequence": moves}
	var doc models.Move
	if err := collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}

	// Construire outcomes pour tous les coups suivants
	outcomes := make(models.Outcome)
	for _, next := range doc.NextMoves {
		outcomes[next.Name] = map[int]uint{
			1:  next.WhiteWin,
			0:  next.Draw,
			-1: next.BlackWin,
		}
	}

	return outcomes, nil
}
