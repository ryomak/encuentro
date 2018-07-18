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
	if err := db.Preload("Users").First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (Event) GetAll(c *gin.Context) (interface{}, error) {
	event := []Event{}
	db := util.GetDB(c)
	if err := db.Preload("Users").Preload("Tags").Preload("Restaurant").Find(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (Event) Delete(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func (Event) Create(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := c.Bind(&event); err != nil {
		return nil, err
	}
	if err := db.Create(&event).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&event).Association("Tags").Append(event.Tags).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&event).Association("Users").Append(event.Users).Error; err != nil {
		return nil, err
	}
	res := Event{}
	if err := db.Preload("Tags").Preload("Users").Preload("Restaurant").First(&res, "id =?", event.ID).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (Event) Update(c *gin.Context) (interface{}, error) {
	event := Event{}
	db := util.GetDB(c)
	if err := c.Bind(&event); err != nil {
		return nil, err
	}
	e := Event{}
	if err := db.First(&e, "id = ?", c.Param("id")).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&e).Association("Users").Replace(event.Users).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&e).Association("Tags").Replace(event.Users).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&e).Updates(&event).Error; err != nil {
		return nil, err
	}
	res := Event{}
	if err := db.Preload("Tags").Preload("Users").Preload("Restaurant").First(&res, "id = ?", c.Param("id")).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
