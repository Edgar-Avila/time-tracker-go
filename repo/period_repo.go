package repo

import (
	"fmt"
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

func (pr *periodRepo) GetAfter(timespan string) []models.Period {
    var results []models.Period
    if timespan == "all" {
        if err := db.Get().Find(&results).Error; err != nil {
            log.Fatal(err)
        }
    } else {
        where := fmt.Sprintf("start_time > date('now', '-1 %ss')", timespan)
        fmt.Println(where)
        if err := db.Get().Where(where).Find(&results).Error; err != nil {
            log.Fatal(err)
        }
    }
    return results
}
