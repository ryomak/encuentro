package repository

import (
	"github.com/jinzhu/gorm"
	"reflect"
)



func GetById(db *gorm.DB,model interface{},id string)(interface{},error){
	m := NewInstance(model)
	err := db.First(m).Where("id =?",id).Error
	return m,err
}

func GetAll(db *gorm.DB,model interface{})(interface{},error){
	m := NewInstance(model)
	err := db.Find(m).Error
	return m,err
}

func GetByOption(db *gorm.DB,model interface{},option ...interface{})(interface{},error){
	m := NewInstance(model)
	err := db.Find(m,option...).Error
	return m,err
}

func GetOneByOption(db *gorm.DB,model interface{},option ...interface{})(interface{},error){
	m := NewInstance(model)
	err := db.First(m,option...).Error
	return m,err
}

func NewInstance(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}