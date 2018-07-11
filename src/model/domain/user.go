package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"github.com/ryomak/encuentro/src/util"
)


type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      uint `json:"age"`
	Gender   uint   `json:"gender"`
	Income   uint   `json:"income"` //年収
	Job      Job `json:"job"`
	LineID   string `json:"line_id"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Tag      []Tag  `json:"tag"`
}

func (User) Migrate(db *gorm.DB)error{
	initData := []User{
		{
			Name:"初期男",
			Age:25,Gender:util.MALE,
			Income:600,
			Job:Job{Name:"エンジニア"},
			LineID:"fds",
			Mail:"testa@a.com",
			Phone:"090",
			Tag:[]Tag{{Name:"年収1000万以上"}},
		},
	}
	for _,v := range initData{
		if err := db.Create(&v).Error;err != nil{
			return err
		}
	}
	return nil
}

func (User)Validate(validate validator.Validate){

}