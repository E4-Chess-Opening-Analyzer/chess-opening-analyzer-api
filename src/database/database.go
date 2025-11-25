package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *gorm.DB

// ConnectMongo connects to MongoDB using environment variables.
// It prefers application credentials (APP_DB_USER/APP_DB_PASSWORD) and
// falls back to DB_USER/DB_PASSWORD if the app creds are not set.
func ConnectMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	if host == "" || port == "" || dbName == "" || user == "" || pass == "" {
		return nil, fmt.Errorf("database environment variables are not properly set")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, pass, host, port, dbName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
