package repo

import (
	"log"
	"time-tracker/db"
	"time-tracker/models"
)

type activityRepo struct {
}

var activityRepoInstance *activityRepo

func ActivityRepo() *activityRepo {
    if activityRepoInstance == nil {
        activityRepoInstance = &activityRepo{}
    }
    return activityRepoInstance
}

// ***************************************
//               Functions
// ***************************************
func (ar *activityRepo) GetAll() []models.Activity {
	var activities []models.Activity
	if err := db.Get().Preload("ActivePeriod").Find(&activities).Error; err != nil {
		log.Fatal(err)
	}
	return activities
}

func (ar *activityRepo) GetByName(name string) models.Activity {
    var activity models.Activity
    if err := db.Get().Where("name = ?", name).Preload("ActivePeriod").First(&activity).Error; err != nil {
        log.Fatal(err)
    }
    return activity
}

func (ar *activityRepo) Create(activity *models.Activity) {
    if err := db.Get().Create(&activity).Error; err != nil {
        log.Fatal(err)
    }
}

func (ar *activityRepo) DeleteByName(name string) {
    if err := db.Get().Unscoped().Delete(&models.Activity{}, "name = ?", name).Error; err != nil {
        log.Fatal(err)
    }
}

func (ar *activityRepo) Update(activity *models.Activity) {
    if err := db.Get().Save(activity).Error; err != nil {
        log.Fatal(err)
    }
}

func (ar *activityRepo) SetFieldNull(activity *models.Activity, field string) {
    if err := db.Get().Model(&activity).Select("active_period_id").Update(field, nil).Error; err != nil {
        log.Fatal(err)
    }
}
