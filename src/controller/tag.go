package controller

import (
	"github.com/ryomak/encuentro/src/model"
	"github.com/ryomak/encuentro/src/util"
)

var TagEndpoint = util.NewHandler(model.Tag{}, &util.ApiHandler{})
