package cmd

import (
	"fmt"
	"log"
	"time"
	"time-tracker/models"
	"time-tracker/repo"
	"time-tracker/util"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "View a report of your activities",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		timespan, err := cmd.Flags().GetString("timespan")
		if err != nil {
			log.Fatal(err)
		}

		activityName, err := cmd.Flags().GetString("activity")
		if err != nil {
			log.Fatal(err)
		}

		// Check if timespan is valid
		validValues := []string{"all", "day", "week", "month", "year"}
		isValid := util.StringInSlice(timespan, validValues)
		if !isValid {
			fmt.Printf("Timespan should be one of %v", validValues)
		}

		// Get results
		var results []models.Record
        if timespan == "all" && activityName == "all" {
            results = repo.RecordRepo().GetAll()
        } else if timespan == "all" && activityName != "all" {
			activity := repo.ActivityRepo().GetByName(activityName)
            results = repo.RecordRepo().GetAllByActivity(activity)
        } else if timespan != "all" && activityName == "all" {
			results = repo.RecordRepo().GetAfter(timespan)
		} else {
			activity := repo.ActivityRepo().GetByName(activityName)
			results = repo.RecordRepo().GetAfterByActivity(timespan, activity)
		}

		for _, result := range results {
			diff := result.EndTime.Sub(result.StartTime)
			startDate := result.StartTime.Format("2006-01-02")
			out := time.Time{}.Add(diff).Format("15:04:05")
			name := result.Activity.Name
			fmt.Printf("%s: Activity %s was done for %s\n", startDate, name, out)
		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().StringP("timespan", "t", "all", "How old should the analytics be?")
	reportCmd.Flags().StringP("activity", "a", "all", "Get reports on a particular activity")
}
