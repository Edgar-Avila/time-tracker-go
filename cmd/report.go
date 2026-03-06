package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/olebedev/when"
	en "github.com/olebedev/when/rules/en"
	"github.com/spf13/cobra"
	myWhen "time-tracker/util/when"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "View a report of your activities",
	Run: func(cmd *cobra.Command, args []string) {
		activityName, err := cmd.Flags().GetString("activity")
		if err != nil {
			log.Fatal(err)
		}
		// Get results
		var results []models.Record

		// If extra args were provided, try parsing them as a natural language time
		if len(args) > 0 {
			text := strings.Join(args, " ")
			w := when.New(nil)
			w.Add(en.All...)
			w.Add(myWhen.All...)
			res, err := w.Parse(text, time.Now())
			if err != nil {
				log.Fatalf("failed to parse time expression: %v", err)
			}
			if res == nil {
				fmt.Printf("Could not understand time expression: %s\n", text)
				return
			}

			fmt.Printf("Reporting on records since %s\n", res.Time.Format(time.RFC1123))
			since := res.Time
			if activityName == "all" {
				results = repo.RecordRepo().GetAfterSince(since)
			} else {
				activity := repo.ActivityRepo().GetByName(activityName)
				results = repo.RecordRepo().GetAfterByActivitySince(since, activity)
			}
		} else {
			if activityName == "all" {
				results = repo.RecordRepo().GetAll()
			} else {
				activity := repo.ActivityRepo().GetByName(activityName)
				results = repo.RecordRepo().GetAllByActivity(activity)
			}
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
	reportCmd.Flags().StringP("activity", "a", "all", "Get reports on a particular activity")
}
