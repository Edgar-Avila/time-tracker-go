package cmd

import (
	"time-tracker/repo"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// list activities
var listActivitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "List all activities",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, activity := range repo.ActivityRepo().GetAll() {
			activeStr := "\n"
			if activity.ActiveRecord != nil {
				activeStr = "(active)\n"
			}
			color.New(color.FgWhite).Printf("- %s %s", activity.Name, activeStr)
		}
	},
}

func init() {
	listCmd.AddCommand(listActivitiesCmd)
}
