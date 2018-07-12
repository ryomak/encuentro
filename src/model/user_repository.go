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

func (User) Delete(c *gin.Context)(interface{}, error)  {
	user := User{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&user).Error; err != nil {
		return nil,err
	}
	return &user,nil
}

func (User) Create(c *gin.Context) (interface{}, error) {
	user := User{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&user).Error; err != nil {
		return user, err
	}
	return &user, nil
}

func (User) Update(c *gin.Context) (interface{}, error)  {
	user := User{}
	db := util.GetDB(c)
	if err := db.Preload("Jobs").Preload("Tags").First(&user).Error; err != nil {
		return  nil,err
	}
	return &user,nil
}
