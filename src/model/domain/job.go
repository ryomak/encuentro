package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Job struct {
	gorm.Model
	Name     string `json:"name"`
}

func (Job) Migrate(db *gorm.DB)error{
	initData := []Job{
		{Name:"エンジニア"},
		{Name:"学生"},
		{Name:"営業"},
	}
	for _,v := range initData{
		if err := db.Create(&v).Error;err != nil{
			return err
		}
	}
	return nil
}

func (Job)Validate(validate validator.Validate){

}