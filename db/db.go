package db

import (
	"log"
	"time-tracker/config"
	"time-tracker/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get() *gorm.DB {
	if db == nil {
		database, err := gorm.Open(sqlite.Open(config.Get().DbPath), &gorm.Config{})
        if err != nil {
            log.Fatal(err)
        }
        db = database
        db.AutoMigrate(&models.Activity{}, &models.Record{})
	}
	return db
}
