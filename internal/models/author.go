package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	PublicName string `gorm:"size:300"`
}
