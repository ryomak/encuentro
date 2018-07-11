package controller

import (
	"github.com/ryomak/encuentro/src/model/domain"
	"github.com/ryomak/encuentro/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)



var UserEndpoint = util.NewHandler(domain.User{},&util.ApiHandler{
	GetHandler: util.GetHandler(domain.User{},util.GetOption{}),
	GetAllHandler:func(c *gin.Context) {
		db := util.GetDB(c)
		u := []domain.User{}
		if err := db.Preload("Jobs").Preload("Tags").Find(&u).Error;
			err != nil{
			util.GetErrorResponse(err)
			return
		}
		c.JSON(http.StatusOK,u)
		return
	},
})

