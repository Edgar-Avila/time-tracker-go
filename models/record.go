package models

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
    StartTime  time.Time  `gorm:"not null"`
	EndTime    time.Time
    ActivityID uint        `gorm:"not null"`
	Activity   *Activity
}
