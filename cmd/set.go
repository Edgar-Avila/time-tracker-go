package cmd

import (
	"fmt"
	"time"
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
// time-tracker set <activity> <start> [end]
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Manually create a record for an activity",
	Long: `Create a record for an activity using explicit start and optional end times.
Time formats accepted: RFC3339 (2006-01-02T15:04:05Z07:00),
"2006-01-02T15:04", or "2006-01-02 15:04". If end is omitted the record will be set as active.
Example: time-tracker set exercise 2023-03-01T09:00:00 2023-03-01T10:30:00`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		activity := repo.ActivityRepo().GetByName(name)

		// parse start time
		parseTime := func(s string) (time.Time, error) {
			// Try full date/time layouts first
			layouts := []string{
				time.RFC3339,
				time.RFC3339Nano,
				"2006-01-02T15:04:05",
				"2006-01-02T15:04",
				"2006-01-02 15:04:05",
				"2006-01-02 15:04",
				"2006-01-02",
			}
			var t time.Time
			var err error
			for _, l := range layouts {
				t, err = time.Parse(l, s)
				if err == nil {
					return t, nil
				}
			}

			// If user provided only a time (e.g. "09:00"), assume today
			timeOnlyLayouts := []string{"15:04:05", "15:04"}
			for _, l := range timeOnlyLayouts {
				pt, err2 := time.Parse(l, s)
				if err2 == nil {
					now := time.Now()
					loc := now.Location()
					return time.Date(now.Year(), now.Month(), now.Day(), pt.Hour(), pt.Minute(), pt.Second(), 0, loc), nil
				}
			}

			return time.Time{}, fmt.Errorf("unrecognized time format: %v", err)
		}

		start, err := parseTime(args[1])
		if err != nil {
			color.New(color.FgRed).Printf("Failed to parse start time: %v\n", err)
			return
		}

		var end time.Time
		hasEnd := false
		if len(args) > 2 {
			end, err = parseTime(args[2])
			if err != nil {
				color.New(color.FgRed).Printf("Failed to parse end time: %v\n", err)
				return
			}
			hasEnd = true
			if end.Before(start) {
				color.New(color.FgRed).Printf("End time is before start time\n")
				return
			}
		}

		record := models.Record{StartTime: start, ActivityID: activity.ID}
		if hasEnd {
			record.EndTime = end
		}

		repo.RecordRepo().Create(&record)

		if !hasEnd {
			// if no end provided, set as active record if activity not already active
			if activity.ActiveRecord != nil {
				color.New(color.FgYellow).Printf("Activity \"%s\" already has an active record\n", name)
			} else {
				activity.ActiveRecordID = &record.ID
				repo.ActivityRepo().Update(&activity)
				color.New(color.FgGreen).Printf("Created active record for activity \"%s\" starting at %s\n", name, start)
			}
		} else {
			color.New(color.FgGreen).Printf("Created record for activity \"%s\" from %s to %s\n", name, start, end)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
