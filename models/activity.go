package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name           string `gorm:"unique"`
	ActivePeriodID *uint
	ActivePeriod   *Period `gorm:"foreignKey:ActivePeriodID"`
}
