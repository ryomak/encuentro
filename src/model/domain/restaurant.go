package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Restaurant struct {
	gorm.Model
	Name     string `json:"name"`
	Address  string `json:"address"`
	Price    uint `json:"price"`
}

func (Restaurant) Migrate(db *gorm.DB)error{
	initData := []Restaurant{
		{Name:"鳥貴族",Address:"大阪府大阪市北区",Price:3000,},
	}
	for _,v := range initData{
		if err := db.Create(&v).Error;err != nil{
			return err
		}
	}
	return nil
}

func (Restaurant)Validate(validate validator.Validate){

}