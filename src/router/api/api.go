package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ryomak/encuentro/src/controller"
	"github.com/ryomak/encuentro/src/util"
)

func Public(n *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	n.Use(middleware...)
	//TODO idは数字のみ
	util.WrapHandler(n.Group("/tag"), controller.TagEndpoint)
	util.WrapHandler(n.Group("/event"), controller.EventEndpoint)
	util.WrapHandler(n.Group("/user"), controller.UserEndpoint)

}
