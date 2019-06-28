package model

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name string
	Age int `gorm:"index;not null"`
}
