package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ryomak/encuentro/src/util"
	"gopkg.in/go-playground/validator.v9"
)

type Event struct {
	gorm.Model
	Name         string     `json:"name"`
	MenNumber    uint       `json:"men_number"`
	WomenNumber  uint       `json:"women_number"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `json:"restaurant" gorm:"foreignkey:RestaurantID;association_autocreate:false;association_autoupdate:false"`
	Tags         []Tag      `json:"tags" gorm:"many2many:event_tags;gorm:association_autocreate:false;association_autoupdate:false"`
	Users        []User     `json:"book" gorm:"many2many:event_users;gorm:association_autocreate:false;association_autoupdate:false"`
}

//もしいるなら
type Tags struct {
	AllTag   []Tag `json:"all_tag"`
	NonTag   []Tag `json:"non_tag"`
	ManTag   []Tag `json:"man_tag"`
	WomanTag []Tag `json:"woman_tag"`
}

func (Event) Migrate(db *gorm.DB) error {
	initData := []Event{
		{
			Name: "エンジニア合コン", MenNumber: 3, WomenNumber: 3,
			Restaurant: Restaurant{
				Name:    "6年４組",
				Address: "大阪駅第６ビル",
				Price:   3000,
			},
			Tags: []Tag{{Name: "CA", Gender: util.FEMALE}},
			Users: []User{
				User{
					Name:   "初期女",
					Age:    28,
					Gender: util.FEMALE,
					Income: 300,
					LineID: "eeee",
					Mail:   "test@a.com",
					Phone:  "090",
				},
			},
		},
	}
	for _, v := range initData {
		if err := db.Create(&v).Error; err != nil {
			return err
		}
	}
	return nil
}

func (Event) Validate(validate validator.Validate) {

}
