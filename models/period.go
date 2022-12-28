package models

import (
	"time"

	"gorm.io/gorm"
)

type Period struct {
	gorm.Model
    StartTime  time.Time  `gorm:"not null"`
	EndTime    time.Time
    ActivityID int        `gorm:"not null"`
	Activity   *Activity
}
