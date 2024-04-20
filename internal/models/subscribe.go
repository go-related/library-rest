package models

import (
	"gorm.io/gorm"
)

type Subscribe struct {
	gorm.Model
	SubscriberID uint     `gorm:""`
	Books        []Book   `gorm:"many2many:subscribe_book;"`
	Authors      []Author `gorm:"many2many:subscribe_author;"`
}
