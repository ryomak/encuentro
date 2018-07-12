package model

import (
	"github.com/gin-gonic/gin"
	"github.com/ryomak/encuentro/src/util"
)

func (User) Get(c *gin.Context) (interface{}, error) {
	user := User{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&user, "id = ?", c.Param("id")).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (User) GetAll(c *gin.Context) (interface{}, error) {
	u := []*User{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (User) Delete(c *gin.Context) (interface{}, error) {
	user := User{}
	db := util.GetDB(c)
	if err := db.First(&user, c.Param("id")).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&user).Association("tags").Clear().Error; err != nil {
		return nil, err
	}
	if err := db.Model(&user).Association("jobs").Clear().Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (User) Create(c *gin.Context) (interface{}, error) {
	user := new(User)
	db := util.GetDB(c)
	if err := c.Bind(&user); err != nil {
		return nil, err
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Jobs").Preload("Tags").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (User) Update(c *gin.Context) (interface{}, error) {
	user := User{}
	db := util.GetDB(c)
	if err := c.Bind(&user); err != nil {
		return nil, err
	}
	u := User{}
	if err := db.First(&u, "id = ?", c.Param("id")).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&u).Updates(&user).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
