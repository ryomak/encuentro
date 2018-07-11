package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func ConnectDB(dbConfig DBConfig)( *gorm.DB){
	param := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?parseTime=true&loc=UTC"
	db, err := gorm.Open("mysql", param)
	if err != nil {
		panic(err)
	}
	if dbConfig.Debug {
		db = db.Debug()
	}
	return db
}

func GetDB(c *gin.Context)*gorm.DB{
	db ,_ :=c.Get(DATABASEKEY)
	return db.(*gorm.DB)
}

