package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/ryomak/encuentro/src/router"
	"github.com/ryomak/encuentro/src/util"
)

var config util.Config

func init() {
	cPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	config = util.LoadConfig(*cPath)
}

func main() {
	gin.SetMode(config.Env)
	router := router.InitRoute(util.ConnectDB(config.Database))
	router.Run(":" + config.Port)
}
