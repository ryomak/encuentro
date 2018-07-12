package util

import "github.com/gin-gonic/gin"

type ApiHandler struct {
	GetHandler gin.HandlerFunc
	GetAllHandler gin.HandlerFunc
	CreateHandler gin.HandlerFunc
	UpdateHandler gin.HandlerFunc
	DeleteHandler gin.HandlerFunc
}

type ApiOption struct {
	GetOption    GetOption
	GetAllOption  GetOption
	CreateOption CreateOption
	UpdateOption UpdateOption
	DeleteOption DeleteOption
}

type CreateOption struct {
	BeforeCreate func (c *gin.Context)error
	AfterCreate  func (c *gin.Context)error
}

type GetOption struct {
	BeforeGet func (c *gin.Context)error
	AfterGet  func (c *gin.Context)error
}


type UpdateOption struct {
	BeforeUpdate func (c *gin.Context)error
	AfterUpdate func (c *gin.Context)error
}

type DeleteOption struct {
	BeforeDelete func (c *gin.Context)error
	AfterDelete func (c *gin.Context)error
}

type Method interface {
	Get(*gin.Context)(interface{},error)
	GetAll(*gin.Context)(interface{},error)
	Create(*gin.Context)(interface{},error)
	Update(*gin.Context)(interface{},error)
	Delete(*gin.Context)(interface{},error)
}

func callExist(c *gin.Context,f func(g *gin.Context)error)error{
	if f != nil{
		return f(c)
	}
	return nil
}