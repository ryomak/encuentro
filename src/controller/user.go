package controller

import (
	"github.com/ryomak/encuentro/src/model"
	"github.com/ryomak/encuentro/src/util"
)

var UserEndpoint = util.NewHandler(model.User{}, &util.ApiHandler{
	GetHandler:    util.GetHandler(model.User{}, util.GetOption{}),
	GetAllHandler: util.GetAllHandler(model.User{}, util.GetOption{}),
	CreateHandler: util.CreateHandler(model.User{}, util.CreateOption{}),
	UpdateHandler: util.UpdateHandler(model.User{}, util.UpdateOption{}),
	DeleteHandler: util.DeleteHandler(model.User{}, util.DeleteOption{}),
})
