package model

import (
	"github.com/gin-gonic/gin"
	"github.com/ryomak/encuentro/src/util"
)

func (Event) Get(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := db.First(&event, c.Param("id")).Related(&event.Restaurant).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Tags").First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (Event) GetAll(c *gin.Context) (interface{}, error) {
	event := []Event{}
	db := util.GetDB(c)
	if err := db.Preload("Tags").Preload("Restaurant").Find(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (Event) Delete(c *gin.Context) (interface{}, error) {
	return nil,nil
}

func (Event) Create(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&event).Error; err != nil {
		return event, err
	}
	return &event, nil
}

func (Event) Update(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
