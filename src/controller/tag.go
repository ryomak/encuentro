package controller

import (
	"github.com/ryomak/encuentro/src/model/domain"
	"github.com/ryomak/encuentro/src/util"
)

var TagEndpoint = util.NewHandler(domain.Tag{},&util.ApiHandler{})

