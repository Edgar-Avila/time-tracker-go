/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"time-tracker/models"
	"time-tracker/repo"

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

		// Find activity
        activity := repo.ActivityRepo().GetByName(name)

		// Check if not doing activiy already
		if activity.ActivePeriod != nil {
			fmt.Printf("You are already doing the activity \"%s\"", name)
		}

		// Create period
		period := models.Period{StartTime: time.Now(), ActivityID: activity.ID}
        repo.PeriodRepo().Create(&period)

		// Link activity with active period
		activity.ActivePeriodID = &period.ID
        repo.ActivityRepo().Update(&activity)

		fmt.Printf("You started the activity \"%s\"\n", name)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
