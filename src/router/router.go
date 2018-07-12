package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ryomak/encuentro/src/router/api"
	"github.com/ryomak/encuentro/src/router/middleware"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	//TODO CSRF対策
	api.Public(r.Group("/api/v1"), middleware.CORSMiddleware(), middleware.TransactionMiddleware(db))

	return r
}
