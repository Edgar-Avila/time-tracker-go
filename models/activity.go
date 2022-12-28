package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
    Name         string  `gorm:"unique"`
	ActivePeriod *Period 
}
