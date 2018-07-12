package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)


type ModelInterface interface {
	Migrate(*gorm.DB)error
	Validate(validator.Validate)
}
