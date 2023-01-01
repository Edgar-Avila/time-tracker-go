package repo

import (
	"fmt"
	"log"
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

func (rr *redordRepo) GetAfter(timespan string) []models.Record {
    var results []models.Record
    where := fmt.Sprintf("start_time > date('now', '-1 %ss')", timespan)
    if err := db.Get().Where(where).Preload("Activity").Find(&results).Error; err != nil {
        log.Fatal(err)
    }
    return results
}

func (rr *redordRepo) GetAfterByActivity(timespan string, activity models.Activity) []models.Record {
    var results []models.Record
    where := fmt.Sprintf("start_time > date('now', '-1 %ss') AND activity_id = ?", timespan)
    if err := db.Get().Where(where, activity.ID).Preload("Activity").Find(&results).Error; err != nil {
        log.Fatal(err)
    }
    return results
}

func (rr *redordRepo) DeleteByActivityId(id uint) {
    if err := db.Get().Unscoped().Delete(&models.Record{}, "activity_id = ?", id).Error; err != nil {
        log.Fatal(err)
    }
}

