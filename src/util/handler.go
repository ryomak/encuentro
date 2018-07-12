package util

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"github.com/ryomak/brand_web/src/util"
	"net/http"
)


func NewHandler(model interface{},handler *ApiHandler)*ApiHandler{
	if handler.GetHandler == nil {
		handler.GetHandler =  getNormalHandler(model,GetOption{})
	}
	if handler.GetAllHandler == nil{
		handler.GetAllHandler = getAllNormalHandler(model,GetOption{})
	}
	if handler.CreateHandler == nil {
		handler.CreateHandler = createNormalHandler(model,CreateOption{})
	}
	if handler.UpdateHandler == nil{
		handler.UpdateHandler = updateNormalHandler(model,UpdateOption{})
	}
	if handler.DeleteHandler == nil {
		handler.DeleteHandler = deleteNormalHandler(model,DeleteOption{})
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

func GetHandler(m Method,op GetOption)gin.HandlerFunc{
	return func(c *gin.Context){
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		res ,err := m.Get(c)
		if err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,res)
		return
	}
}

func GetAllHandler(m Method,op GetOption)gin.HandlerFunc{
	return func(c *gin.Context){
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		res, err := m.GetAll(c)
		if err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,res)
		return
	}
}


func CreateHandler(m Method,op CreateOption)gin.HandlerFunc{
	return func(c *gin.Context){
		if err := callExist(c,op.BeforeCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		res ,err := m.Create(c)
		if err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,res)
		return
	}
}

func UpdateHandler(m Method,op UpdateOption)gin.HandlerFunc{
	return func(c *gin.Context){
		if err := callExist(c,op.BeforeUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		res , err := m.Update(c)
		if err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,res)
		return
	}
}


func DeleteHandler(m Method,op DeleteOption)gin.HandlerFunc{
	return func(c *gin.Context){
		if err := callExist(c,op.BeforeDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		res, err := m.Delete(c)
		if err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,res)
		return
	}
}

func getNormalHandler(m interface{},op GetOption)gin.HandlerFunc{
	return func(c *gin.Context) {
		p := NewInstance(m)
		db := util.GetDB(c)
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		if err := db.First(p,"id = ?",c.Param("id")).Error;err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func getAllNormalHandler(m interface{},op GetOption)gin.HandlerFunc{
	p := NewSlice(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := callExist(c,op.BeforeGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		if err := db.Find(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterGet);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func createNormalHandler(m interface{},op CreateOption)gin.HandlerFunc{
	p := NewInstance(m)
	return func(c *gin.Context) {
		db := util.GetDB(c)
		if err := c.Bind(p);err != nil{
			c.JSON(http.StatusBadRequest,CreateErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.BeforeCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		if err := db.Create(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,DBCreateErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterCreate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func updateNormalHandler(m interface{},op UpdateOption)gin.HandlerFunc{
	return func(c *gin.Context) {
		p := NewInstance(m)
		updateP := NewInstance(m)
		db := util.GetDB(c)
		if err := c.Bind(updateP);err != nil{
			c.JSON(http.StatusBadRequest,CreateErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.BeforeUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		if err := db.Model(&p).Update(updateP).Error;err != nil{
			c.JSON(http.StatusBadRequest,DBCreateErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterUpdate);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		c.JSON(http.StatusOK,p)
		return
	}
}

func deleteNormalHandler(m interface{},op DeleteOption)gin.HandlerFunc{
	return func(c *gin.Context) {
		p := NewInstance(m)
		db := util.GetDB(c)
		if err := db.First(p,"id = ?",c.Param("id")).Error;err!= nil{
			c.JSON(http.StatusBadRequest,GetErrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.BeforeDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
			return
		}
		if err := db.Delete(p).Error;err != nil{
			c.JSON(http.StatusBadRequest,DeleterrorResponse(err))
			panic(err)
			return
		}
		if err := callExist(c,op.AfterDelete);err != nil{
			c.JSON(http.StatusBadRequest,NomalErrorResponse(err))
			panic(err)
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

// util.NewInstance returns reference of new slice of i
func NewSlice(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	ptr := reflect.New(slice.Type())
	ptr.Elem().Set(slice)
	return ptr.Interface()
}
