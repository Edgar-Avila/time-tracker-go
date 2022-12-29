/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"
	"time-tracker/db"
	"time-tracker/models"

	"github.com/spf13/cobra"
)

// endCmd represents the end command
// time-tracker end exercise
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End an activity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var activity models.Activity

		// Get activity
		if err := db.Get().Where("name = ?", name).Preload("ActivePeriod").First(&activity).Error; err != nil {
			log.Fatal(err)
		}

		// Check if activity was not active right now
		if activity.ActivePeriod == nil {
			fmt.Printf("You are not doing the activity \"%s\"", name)
			return
		}

		// Update active period with end time
		activity.ActivePeriod.EndTime = time.Now()
		if err := db.Get().Save(&activity.ActivePeriod).Error; err != nil {
			log.Fatal(err)
		}

		// Remove active period from activity since it is already finished
		activity.ActivePeriodID = nil
		if err := db.Get().Model(&activity).Select("active_period_id").Updates(activity).Error; err != nil {
			log.Fatal(err)
		}

		fmt.Printf("You finished doing the activity \"%s\"", name)
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
