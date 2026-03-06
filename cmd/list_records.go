package cmd

import (
	"fmt"
	"time"
	"time-tracker/repo"

	"github.com/spf13/cobra"
)

// list records [activityName]
var listRecordsCmd = &cobra.Command{
	Use:   "records [activityName]",
	Short: "List records, optionally filtered by activity",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			activity := repo.ActivityRepo().GetByName(args[0])
			for _, r := range repo.RecordRepo().GetAllByActivity(activity) {
				endStr := "(active)"
				if !r.EndTime.IsZero() {
					endStr = r.EndTime.Format(time.RFC3339)
				}
				fmt.Printf("- %s: %s -> %s\n", r.Activity.Name, r.StartTime.Format(time.RFC3339), endStr)
			}
			return
		}

		for _, r := range repo.RecordRepo().GetAll() {
			endStr := "(active)"
			if !r.EndTime.IsZero() {
				endStr = r.EndTime.Format(time.RFC3339)
			}
			name := "<unknown>"
			if r.Activity != nil {
				name = r.Activity.Name
			}
			fmt.Printf("- %s: %s -> %s\n", name, r.StartTime.Format(time.RFC3339), endStr)
		}
	},
}

func init() {
	listCmd.AddCommand(listRecordsCmd)
}
