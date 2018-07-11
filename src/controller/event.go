package controller

import (
	"github.com/ryomak/encuentro/src/model/domain"
	"github.com/ryomak/encuentro/src/util"
	"github.com/gin-gonic/gin"
)

var EventEndpoint = util.NewHandler(domain.Event{},&util.ApiHandler{
	GetHandler: func(c *gin.Context) {

	},
})
