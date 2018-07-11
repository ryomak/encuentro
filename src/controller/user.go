package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ryomak/encuentro/src/model/domain"
	"github.com/ryomak/encuentro/src/util"
	"github.com/ryomak/encuentro/src/model/repository"
)

func GetAllUserHandler(c *gin.Context) {
	db := util.GetDB(c)
	users, err := repository.GetByOption(db, domain.User{},"gender = ?",1)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK,users)
}

func GetUserHandler(c *gin.Context) {
	db := util.GetDB(c)
	user, err := repository.GetById(db, domain.User{},c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GetErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK,user)
}