package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"
	"time-tracker/repo"

	myWhen "time-tracker/util/when"

	"github.com/spf13/cobra"
)

// list records
var listRecordsCmd = &cobra.Command{
	Use:   "records",
	Short: "List records, optionally filtered by activity",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activityName, err := cmd.Flags().GetString("activity")
		if err != nil {
			log.Fatal(err)
		}

		// If args provided, parse them as a natural language time expression
		if len(args) > 0 {
			text := strings.Join(args, " ")
			w := myWhen.New()
			res, err := w.Parse(text, time.Now())
			if err != nil {
				log.Fatalf("failed to parse time expression: %v", err)
			}
			if res == nil {
				fmt.Printf("Could not understand time expression: %s\n", text)
				return
			}
			since := res.Time
			if activityName == "all" {
				fmt.Printf("Listing records since %s\n", since.Format(time.RFC1123))
				for _, r := range repo.RecordRepo().GetAfterSince(since) {
					endStr := "(active)"
					if !r.EndTime.IsZero() {
						endStr = r.EndTime.Format(time.RFC3339)
					}
					name := "<unknown>"
					if r.Activity != nil {
						name = r.Activity.Name
					}
					fmt.Printf("- %d: %s: %s -> %s\n", r.ID, name, r.StartTime.Format(time.RFC3339), endStr)
				}
			} else {
				activity := repo.ActivityRepo().GetByName(activityName)
				fmt.Printf("Listing records for activity %s since %s\n", activity.Name, since.Format(time.RFC1123))
				for _, r := range repo.RecordRepo().GetAfterByActivitySince(since, activity) {
					endStr := "(active)"
					if !r.EndTime.IsZero() {
						endStr = r.EndTime.Format(time.RFC3339)
					}
					fmt.Printf("- %d: %s: %s -> %s\n", r.ID, r.Activity.Name, r.StartTime.Format(time.RFC3339), endStr)
				}
			}
			return
		}

		// No time expression: list either all or by activity
		if activityName == "all" {
			for _, r := range repo.RecordRepo().GetAll() {
				endStr := "(active)"
				if !r.EndTime.IsZero() {
					endStr = r.EndTime.Format(time.RFC3339)
				}
				name := "<unknown>"
				if r.Activity != nil {
					name = r.Activity.Name
				}
				fmt.Printf("- %d: %s: %s -> %s\n", r.ID, name, r.StartTime.Format(time.RFC3339), endStr)
			}
			return
		}

		activity := repo.ActivityRepo().GetByName(activityName)
		for _, r := range repo.RecordRepo().GetAllByActivity(activity) {
			endStr := "(active)"
			if !r.EndTime.IsZero() {
				endStr = r.EndTime.Format(time.RFC3339)
			}
			fmt.Printf("- %d: %s: %s -> %s\n", r.ID, r.Activity.Name, r.StartTime.Format(time.RFC3339), endStr)
		}
	},
}

func init() {
	listCmd.AddCommand(listRecordsCmd)
	listRecordsCmd.Flags().StringP("activity", "a", "all", "Get records for a particular activity (or 'all')")
}
