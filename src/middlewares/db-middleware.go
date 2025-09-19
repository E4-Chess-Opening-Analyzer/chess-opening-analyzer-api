package middlewares

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("DB", db)
        c.Next()
    }
}
