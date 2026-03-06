package cmd

import (
	"fmt"
	"time-tracker/repo"

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
			fmt.Printf("- %s %s", activity.Name, activeStr)
		}
	},
}

func init() {
	listCmd.AddCommand(listActivitiesCmd)
}
