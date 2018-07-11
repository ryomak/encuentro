package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ryomak/brand_web/src/util"
	"github.com/ryomak/brand_web/src/model/repository"
)

func CategoryEndpoint(c *gin.Context) {
	db := util.GetDB(c)
	categories ,err := repository.GetCategoryById(db,c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest,util.MessageResponse{
			Message:"データ取得失敗",
			Detail:err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,categories)
	return
}

func AllCategoryEndpoint(c *gin.Context) {
	db := util.GetDB(c)
	categories ,err := repository.GetCategories(db)
	if err != nil{
		c.JSON(http.StatusBadRequest,util.MessageResponse{
			Message:"データ取得失敗",
			Detail:err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,categories)
	return
}

func CategoryDetailEndpoint(c *gin.Context){
	db := util.GetDB(c)
	detailCategories ,err := repository.GetDetailCategoryById(db,c.Param("id"),c.Param("detail_id"))
	if err != nil{
		c.JSON(http.StatusBadRequest,util.MessageResponse{
			Message:"データ取得失敗",
			Detail:err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,detailCategories)
	return
}

func AllCategoryDetailsEndpoint(c *gin.Context){
	db := util.GetDB(c)
	detailCategories ,err := repository.GetDetailCategoies(db,c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest,util.MessageResponse{
			Message:"データ取得失敗",
			Detail:err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,detailCategories)
	return
}