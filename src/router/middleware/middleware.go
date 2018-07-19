package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ryomak/brand_web/src/util"
)

func TransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin()
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
			} else {
				tx.Commit()

			}
		}()
		c.Set(util.DATABASEKEY, tx)

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthJWTMiddleware(db *gorm.DB) gin.HandlerFunc {

	return
}
