package cmd

import (
	"log"
	"strings"
	"time"
	"time-tracker/repo"

	myWhen "time-tracker/util/when"

	"github.com/fatih/color"
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
				color.New(color.FgRed).Fprintf(color.Output, "failed to parse time expression: %v\n", err)
				return
			}
			if res == nil {
				color.New(color.FgYellow).Printf("Could not understand time expression: %s\n", text)
				return
			}
			since := res.Time
			if activityName == "all" {
				color.New(color.FgGreen).Printf("Listing records since %s\n", since.Format(time.RFC1123))
				for _, r := range repo.RecordRepo().GetAfterSince(since) {
					endStr := "(active)"
					if !r.EndTime.IsZero() {
						endStr = r.EndTime.Format(time.RFC3339)
					}
					name := "<unknown>"
					if r.Activity != nil {
						name = r.Activity.Name
					}
					color.New(color.FgWhite).Printf("- %d: %s: %s -> %s\n", r.ID, name, r.StartTime.Format(time.RFC3339), endStr)
				}
			} else {
				activity := repo.ActivityRepo().GetByName(activityName)
				color.New(color.FgGreen).Printf("Listing records for activity %s since %s\n", activity.Name, since.Format(time.RFC1123))
				for _, r := range repo.RecordRepo().GetAfterByActivitySince(since, activity) {
					endStr := "(active)"
					if !r.EndTime.IsZero() {
						endStr = r.EndTime.Format(time.RFC3339)
					}
					color.New(color.FgWhite).Printf("- %d: %s: %s -> %s\n", r.ID, r.Activity.Name, r.StartTime.Format(time.RFC3339), endStr)
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
				color.New(color.FgWhite).Printf("- %d: %s: %s -> %s\n", r.ID, name, r.StartTime.Format(time.RFC3339), endStr)
			}
			return
		}

		activity := repo.ActivityRepo().GetByName(activityName)
		for _, r := range repo.RecordRepo().GetAllByActivity(activity) {
			endStr := "(active)"
			if !r.EndTime.IsZero() {
				endStr = r.EndTime.Format(time.RFC3339)
			}
			color.New(color.FgWhite).Printf("- %d: %s: %s -> %s\n", r.ID, r.Activity.Name, r.StartTime.Format(time.RFC3339), endStr)
		}
	},
}

func init() {
	listCmd.AddCommand(listRecordsCmd)
	listRecordsCmd.Flags().StringP("activity", "a", "all", "Get records for a particular activity (or 'all')")
}
