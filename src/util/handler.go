package util

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"github.com/ryomak/brand_web/src/util"
	"net/http"
)

func NewHandler(model interface{},handler *ApiHandler)*ApiHandler{
	if handler.GetHandler == nil {
		handler.GetAllHandler =  getHandler(model,GetOption{})
	}
	if handler.GetAllHandler == nil{
		handler.GetAllHandler = getAllHandler(model,GetOption{})
	}
	if handler.CreateHandler == nil {
		handler.CreateHandler = createHandler(model,CreateOption{})
	}
	if handler.UpdateHandler == nil{
		handler.UpdateHandler = updateHandler(model,UpdateOption{})
	}
	if handler.DeleteHandler == nil {
		handler.DeleteHandler = deleteHandler(model,DeleteOption{})
	}
	return handler
}

func WrapHandler(g *gin.RouterGroup,h *ApiHandler,m ...gin.HandlerFunc){
	g.Use(m...)
	g.GET("/:id",h.GetHandler)
	g.GET("",h.GetAllHandler)
	g.POST("",h.CreateHandler)
	g.PUT("/:id",h.UpdateHandler)
	g.DELETE("/:id",h.DeleteHandler)
}

func getHandler(m interface{},op GetOption)gin.HandlerFunc{
	p := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		if err := db.First(p,"id = ?",c.Param("id")).Error;err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func getAllHandler(m interface{},op GetOption)gin.HandlerFunc{
	p := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		if err := db.Find(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func createHandler(m interface{},op CreateOption)gin.HandlerFunc{
	p := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := c.Bind(p);err != nil{
			c.JSON(http.StatusBadRequest,CreateErrorResponse(err))
			return
		}
		if err := callExist(c,op.BeforeCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
		}
		if err := db.Create(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,DBCreateErrorResponse(err))
			return
		}
		if err := callExist(c,op.AfterCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func updateHandler(m interface{},op UpdateOption)gin.HandlerFunc{
	p := NewInstance(m)
	updateP := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := c.Bind(updateP);err != nil{
			c.JSON(http.StatusBadRequest,CreateErrorResponse(err))
			return
		}
		if err := callExist(c,op.BeforeUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		if err := db.Model(&p).Update(updateP).Error;err != nil{
			c.JSON(http.StatusBadRequest,DBCreateErrorResponse(err))
			return
		}
		if err := callExist(c,op.AfterUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func deleteHandler(m interface{},op DeleteOption)gin.HandlerFunc{
	p := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := db.First(p,"id = ?",c.Param("id")).Error;err!= nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			return
		}
		if err := callExist(c,op.BeforeDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		if err := db.Delete(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,DeleterrorResponse(err))
			return
		}
		if err := callExist(c,op.AfterDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func NewInstance(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}
