package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoMiddleware(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func GetMongoDB(c *gin.Context) *mongo.Database {
	db, exists := c.Get("DB")
	if !exists {
		return nil
	}
	return db.(*mongo.Database)
}