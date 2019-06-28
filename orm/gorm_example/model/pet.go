package model

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Name     string
	Person   Person
	PersonID uint
}