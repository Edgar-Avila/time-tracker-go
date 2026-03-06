package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
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
			w := myWhen.New()
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

		// If export flag provided, write results to an Excel file
		exportFile, err := cmd.Flags().GetString("export")
		if err != nil {
			log.Fatal(err)
		}
		if exportFile != "" {
			f := excelize.NewFile()
			sheet := f.GetSheetName(0)
			// headers
			f.SetCellValue(sheet, "A1", "Date")
			f.SetCellValue(sheet, "B1", "Activity")
			f.SetCellValue(sheet, "C1", "Duration")

			for i, result := range results {
				row := i + 2
				diff := result.EndTime.Sub(result.StartTime)
				if result.EndTime.IsZero() {
					diff = time.Since(result.StartTime)
				}
				startDate := result.StartTime.Format("2006-01-02")
				out := time.Time{}.Add(diff).Format("15:04:05")
				name := ""
				if result.Activity != nil {
					name = result.Activity.Name
				}
				f.SetCellValue(sheet, fmt.Sprintf("A%d", row), startDate)
				f.SetCellValue(sheet, fmt.Sprintf("B%d", row), name)
				f.SetCellValue(sheet, fmt.Sprintf("C%d", row), out)
			}
			if err := f.SaveAs(exportFile); err != nil {
				log.Fatalf("failed to save excel file: %v", err)
			}
			fmt.Printf("Exported report to %s\n", exportFile)
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
	reportCmd.Flags().StringP("export", "e", "", "Write report to an Excel file")
}
