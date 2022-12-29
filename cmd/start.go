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

// startCmd represents the start command
// time-tracker start exercise
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an activity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var activity models.Activity

		// Find activity
		if err := db.Get().Where("name = ?", name).Preload("ActivePeriod").First(&activity).Error; err != nil {
			log.Fatal(err)
		}

        // Check if not doing activiy already
        if activity.ActivePeriod != nil {
            fmt.Printf("You are already doing the activity \"%s\"", name)
            return
        }

		// Create period
		period := models.Period{StartTime: time.Now(), ActivityID: activity.ID}
		if err := db.Get().Create(&period).Error; err != nil {
			log.Fatal(err)
		}
        fmt.Printf("You started the activity \"%s\"\n", name)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
