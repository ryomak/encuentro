package domain

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Event struct {
	gorm.Model
	Name     string `json:"name"`
	MenNumber   uint  `json:"men_number"`
	WomenNumber uint `json:"women_number"`
	Restaurant Restaurant `json:"restaurant"`
	Tags       Tags `json:"tags"`
}

type Tags struct {
	AllTag []Tag `json:"all_tag"`
	ManTag []Tag `json:"man_tag"`
	WomanTag []Tag `json:"woman_tag"`
}

func (Event) Migrate(db *gorm.DB)error{
	initData := []Event{
		{Name:"エンジニア合コン",MenNumber:3,WomenNumber:3,Restaurant:Restaurant{
			Name:"6年４組",
			Address:"大阪駅第６ビル",
			Price:3000,
		},
		Tags:Tags{
			[]Tag{{Name:"30歳以下"}},
			[]Tag{{Name:"年収600万以上"}},
			[]Tag{{Name:"エンジニア"}},
		}},
	}
	for _,v := range initData{
		if err := db.Create(&v).Error;err != nil{
			return err
		}
	}
	return nil
}

func (Event)Validate(validate validator.Validate){

}