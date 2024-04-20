package models

import "gorm.io/gorm"

type Subscriber struct {
	gorm.Model
	Email    string `gorm:"size:200"`
	FullName string `gorm:"size:500"`
}
