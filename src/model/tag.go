package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"github.com/ryomak/encuentro/src/util"
)

type Tag struct {
	gorm.Model
	Name     string `json:"name"`
	Gender   uint   `json:"gender"`
}

func (Tag) Migrate(db *gorm.DB)error{
	initData := []Tag{
		{Name:"エンジニア",Gender:util.MALE},
		{Name:"年収600万以上",Gender:util.FEMALE},
		{Name:"大手",Gender:util.BOTH},
	}
	for _,v := range initData{
		if err := db.Create(&v).Error;err != nil{
			return err
		}
	}
	return nil
}

func (Tag)Validate(validate validator.Validate){

}