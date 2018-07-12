package main

import (
	"flag"
	"github.com/ryomak/brand_web/src/util"
	"github.com/ryomak/encuentro/src/model"
	"github.com/sirupsen/logrus"
)

var config util.Config

func init() {
	cPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	config = util.LoadConfig(*cPath)
}

func main() {
	tables := []model.ModelInterface{
		model.Event{},
		model.Job{},
		model.Restaurant{},
		model.Tag{},
		model.User{},
	}
	db := util.ConnectDB(config.Database)
	for _, v := range tables {
		if err := db.AutoMigrate(v).Error; err != nil {
			logrus.Error(err)
		}
	}
	for _, v := range tables {
		if err := v.Migrate(db); err != nil {
			logrus.Error(err)
		}
	}
}
