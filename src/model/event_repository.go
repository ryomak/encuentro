package model

import (
	"github.com/gin-gonic/gin"
	"github.com/ryomak/encuentro/src/util"
)

func (Event) Get(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := db.First(&event,c.Param("id")).Related(&event.Restaurant).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Tags").First(&event).Error;err != nil{
		return nil, err
	}
	tags := Tags{}
	for _,tag := range event.Tags{
		switch tag.Gender {
		case util.None:tags.NonTag = append(tags.NonTag,tag)
		case util.MALE:tags.ManTag = append(tags.ManTag,tag)
		case util.FEMALE:tags.ManTag = append(tags.WomanTag,tag)
		case util.BOTH:tags.ManTag = append(tags.AllTag,tag)
		}
	}
	res := ResEvent{
		Name:event.Name,
		Number:&Num{Man:event.MenNumber,Woman:event.WomenNumber,},
		Restaurant:&event.Restaurant,
		Tags:&tags,
	}
	return &res, nil
}

func (Event) GetAll(c *gin.Context) (interface{}, error) {
	event := []Event{}
	db := util.GetDB(c)
	if err := db.Preload("Restaurant").Find(&event,c.Param("id")).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Tags").First(&event).Error;err != nil{
		return nil, err
	}
	return &event, nil
}

func (Event) Delete(c *gin.Context)(interface{}, error)  {
	event := Event{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&event).Error; err != nil {
		return nil,err
	}
	return &event,nil
}

func (Event) Create(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&event).Error; err != nil {
		return event, err
	}
	return &event, nil
}

func (Event) Update(c *gin.Context) (interface{}, error)  {
	event := Event{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&event).Error; err != nil {
		return  nil,err
	}
	return &event,nil
}