package main
import (
	"github.com/ryomak/brand_web/src/util"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/ryomak/encuentro/src/model/domain"
)

var config util.Config

func init(){
	cPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	config = util.LoadConfig(*cPath)
}

func main() {
	tables := []domain.ModelInterface{
		domain.Event{},
		domain.Job{},
		domain.Restaurant{},
		domain.Tag{},
		domain.User{},
	}
	db := util.ConnectDB(config.Database)
	for _,v := range tables{
		if err := db.AutoMigrate(v).Error;err !=nil{
			logrus.Error(err)
		}
	}
	for _,v := range tables {
		if err := v.Migrate(db);err != nil{
			logrus.Error(err)
		}
	}
}