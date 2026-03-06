package cmd

import (
	"fmt"
	"time-tracker/repo"

	"github.com/spf13/cobra"
)

// remove activity by name
var removeActivityCmd = &cobra.Command{
	Use:   "activity [name]",
	Short: "Remove an activity and its records by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		activity := repo.ActivityRepo().GetByName(name)
		repo.ActivityRepo().DeleteByName(name)
		repo.RecordRepo().DeleteByActivityId(activity.ID)
		fmt.Printf("You removed the activity \"%s\"\n", name)
	},
}

func init() {
	removeCmd.AddCommand(removeActivityCmd)
}
