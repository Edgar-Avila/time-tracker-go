package repo

import (
	"log"
	"time-tracker/db"
	"time-tracker/models"
)

type periodRepo struct {
}

var periodRepoInstance *periodRepo

func PeriodRepo() *periodRepo {
	if periodRepoInstance == nil {
		periodRepoInstance = &periodRepo{}
	}
	return periodRepoInstance
}

// ***************************************
//               Functions
// ***************************************

func (pr *periodRepo) Create(period *models.Period) {
    if err := db.Get().Create(period).Error; err != nil {
        log.Fatal(err)
    }
}

func (pr *periodRepo) Update(period *models.Period) {
    if err := db.Get().Save(period).Error; err != nil {
        log.Fatal(err)
    }
}
