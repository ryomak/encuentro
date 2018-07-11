package util

import "github.com/gin-gonic/gin"

type ApiHandler struct {
	GetHandler gin.HandlerFunc
	GetAllHandler gin.HandlerFunc
	CreateHandler gin.HandlerFunc
	UpdateHandler gin.HandlerFunc
	DeleteHandler gin.HandlerFunc
}

type CreateHandler struct {
	Handler gin.HandlerFunc
	Option CreateOption
}

type CreateOption struct {
	BeforeCreate func (c *gin.Context)error
	AfterCreate  func (c *gin.Context)error
}

type GetHandler struct {
	Handler gin.HandlerFunc
	Option GetOption
}

type GetOption struct {
	BeforeGet func (c *gin.Context)error
	AfterGet  func (c *gin.Context)error
}

type UpdateHandler struct {
	Handler gin.HandlerFunc
	Option UpdateOption
}

type UpdateOption struct {
	BeforeUpdate func (c *gin.Context)error
	AfterUpdate func (c *gin.Context)error
}

type DeleteHandler struct {
	Handler gin.HandlerFunc
	Option DeleteOption
}

type DeleteOption struct {
	BeforeDelete func (c *gin.Context)error
	AfterDelete func (c *gin.Context)error
}

func callExist(c *gin.Context,f func(g *gin.Context)error)error{
	if f != nil{
		return f(c)
	}
	return nil
}