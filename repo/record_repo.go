package repo

import (
	"log"
	"time"
	"time-tracker/db"
	"time-tracker/models"
)

type redordRepo struct {
}

var recordRepoInstance *redordRepo

func RecordRepo() *redordRepo {
	if recordRepoInstance == nil {
		recordRepoInstance = &redordRepo{}
	}
	return recordRepoInstance
}

// ***************************************
//               Functions
// ***************************************

func (rr *redordRepo) Create(record *models.Record) {
	if err := db.Get().Create(record).Error; err != nil {
		log.Fatal(err)
	}
}

func (rr *redordRepo) Update(record *models.Record) {
	if err := db.Get().Save(record).Error; err != nil {
		log.Fatal(err)
	}
}

func (rr *redordRepo) GetAll() []models.Record {
	var results []models.Record
	if err := db.Get().Preload("Activity").Find(&results).Error; err != nil {
		log.Fatal(err)
	}
	return results
}

func (rr *redordRepo) GetAllByActivity(activity models.Activity) []models.Record {
	var results []models.Record
	if err := db.Get().Where("activity_id = ?", activity.ID).Preload("Activity").Find(&results).Error; err != nil {
		log.Fatal(err)
	}
	return results
}

// GetAfterSince returns all records that started after the provided time
func (rr *redordRepo) GetAfterSince(since time.Time) []models.Record {
	var results []models.Record
	if err := db.Get().Where("start_time > ?", since).Preload("Activity").Find(&results).Error; err != nil {
		log.Fatal(err)
	}
	return results
}

// GetAfterByActivitySince returns all records for an activity that started after the provided time
func (rr *redordRepo) GetAfterByActivitySince(since time.Time, activity models.Activity) []models.Record {
	var results []models.Record
	if err := db.Get().Where("start_time > ? AND activity_id = ?", since, activity.ID).Preload("Activity").Find(&results).Error; err != nil {
		log.Fatal(err)
	}
	return results
}

func (rr *redordRepo) DeleteByActivityId(id uint) {
	if err := db.Get().Unscoped().Delete(&models.Record{}, "activity_id = ?", id).Error; err != nil {
		log.Fatal(err)
	}
}

func (rr *redordRepo) DeleteByID(id uint) error {
	if err := db.Get().Unscoped().Delete(&models.Record{}, id).Error; err != nil {
		return err
	}
	return nil
}
