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
	Run: func(cmd *cobra.Command, args []string) {
        for _, name := range args {
            // Find activity
            activity := repo.ActivityRepo().GetByName(name)

            // Check if not doing activiy already
            if activity.ActiveRecord != nil {
                fmt.Printf("You are already doing the activity \"%s\"\n", name)
                continue
            }

            // Create record
            record := models.Record{StartTime: time.Now(), ActivityID: activity.ID}
            repo.RecordRepo().Create(&record)

            // Link activity with active record
            activity.ActiveRecordID = &record.ID
            repo.ActivityRepo().Update(&activity)

            fmt.Printf("You started the activity \"%s\"\n", name)
        }
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
