package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ryomak/encuentro/src/model/domain"
	"github.com/ryomak/encuentro/src/util"
	"github.com/ryomak/encuentro/src/model/repository"
)

var TagEndpoint = util.NewHandler(domain.Tag{},&util.ApiHandler{})

