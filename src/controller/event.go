package controller

import (
	"github.com/ryomak/encuentro/src/model"
	"github.com/ryomak/encuentro/src/util"
)

var EventEndpoint = util.NewHandler(model.Event{}, &util.ApiHandler{
	GetHandler:    util.GetHandler(model.Event{}, util.GetOption{}),
	GetAllHandler: util.GetAllHandler(model.Event{}, util.GetOption{}),
})
