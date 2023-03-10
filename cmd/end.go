package cmd

import (
	"fmt"
	"time"
	"time-tracker/repo"

	"github.com/spf13/cobra"
)

// endCmd represents the end command
// time-tracker end exercise
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End an activity",
	Run: func(cmd *cobra.Command, args []string) {
        for _, name := range args {
            activity := repo.ActivityRepo().GetByName(name)

            // Check if activity was not active right now
            if activity.ActiveRecord == nil {
                fmt.Printf("You are not doing the activity \"%s\"\n", name)
                continue
            }

            // Update active record with end time
            activity.ActiveRecord.EndTime = time.Now()
            repo.RecordRepo().Update(activity.ActiveRecord)

            // Remove active record from activity since it is already finished
            repo.ActivityRepo().SetFieldNull(&activity, "active_record_id")

            fmt.Printf("You finished doing the activity \"%s\"\n", name)
        }
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
