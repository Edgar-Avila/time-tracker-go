package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name           string `gorm:"unique"`
	ActiveRecordID *uint
	ActiveRecord   *Record `gorm:"foreignKey:ActiveRecordID"`
}
