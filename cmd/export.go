package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	myWhen "time-tracker/util/when"
)

// exportCmd writes a report to an Excel file
var exportCmd = &cobra.Command{
	Use:   "export <file> [time expression]",
	Short: "Export records to an Excel file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		activityName, err := cmd.Flags().GetString("activity")
		if err != nil {
			color.New(color.FgRed).Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		// first arg is output file
		exportFile := args[0]

		var results []models.Record

		// If there are additional args, parse them as a natural language time expression
		if len(args) > 1 {
			text := strings.Join(args[1:], " ")
			w := myWhen.New()
			res, err := w.Parse(text, time.Now())
			if err != nil {
				color.New(color.FgRed).Fprintf(os.Stderr, "failed to parse time expression: %v\n", err)
				os.Exit(1)
			}
			if res == nil {
				color.New(color.FgYellow).Printf("Could not understand time expression: %s\n", text)
				return
			}
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

		// Write results to Excel file
		f := excelize.NewFile()
		sheet := f.GetSheetName(0)
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
			// write date and activity
			name := ""
			if result.Activity != nil {
				name = result.Activity.Name
			}
			f.SetCellValue(sheet, fmt.Sprintf("A%d", row), startDate)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", row), name)

			// write duration as Excel time (fraction of a day) so Excel can sum it
			durDays := diff.Seconds() / 86400.0
			ccell := fmt.Sprintf("C%d", row)
			if err := f.SetCellValue(sheet, ccell, durDays); err != nil {
				color.New(color.FgRed).Fprintf(os.Stderr, "failed to set cell value: %v\n", err)
				os.Exit(1)
			}
		}

		// apply time style to duration column and add total formula
		lastRow := len(results) + 1
		// create a time style (use a built-in/custom number format index)
		timeStyle, err := f.NewStyle(&excelize.Style{NumFmt: 46})
		if err == nil {
			// apply style to each duration cell
			if lastRow >= 2 {
				start := fmt.Sprintf("C2")
				end := fmt.Sprintf("C%d", lastRow)
				if err := f.SetCellStyle(sheet, start, end, timeStyle); err != nil {
					color.New(color.FgYellow).Fprintf(os.Stderr, "warning: failed to set cell style: %v\n", err)
				}
				// write total label and formula
				totalCell := fmt.Sprintf("C%d", lastRow+1)
				f.SetCellValue(sheet, fmt.Sprintf("B%d", lastRow+1), "Total")
				f.SetCellFormula(sheet, totalCell, fmt.Sprintf("SUM(%s:%s)", start, end))
				// apply style to total cell
				if err := f.SetCellStyle(sheet, totalCell, totalCell, timeStyle); err != nil {
					color.New(color.FgYellow).Fprintf(os.Stderr, "warning: failed to set total cell style: %v\n", err)
				}
			}
		} else {
			color.New(color.FgYellow).Fprintf(os.Stderr, "warning: failed to create time style: %v\n", err)
		}

		if err := f.SaveAs(exportFile); err != nil {
			color.New(color.FgRed).Fprintf(os.Stderr, "failed to save excel file: %v\n", err)
			os.Exit(1)
		}
		color.New(color.FgGreen).Printf("Exported report to %s\n", exportFile)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("activity", "a", "all", "Get records for a particular activity (or 'all')")
}
