package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `gorm:"size:255"`
}
