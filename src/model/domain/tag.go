package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Tag struct {
	gorm.Model
	Name     string `json:"name"`
}

func (Tag) Migrate(db *gorm.DB)error{
	initData := []Tag{
		{Name:"エンジニア"},
		{Name:"年収600万以上"},
		{Name:"大手"},
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