package cmd

import (
	"fmt"
	"time-tracker/repo"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an activity",
	Run: func(cmd *cobra.Command, args []string) {
        for _, name := range args {
            activity := repo.ActivityRepo().GetByName(name)
            repo.ActivityRepo().DeleteByName(name)
            repo.RecordRepo().DeleteByActivityId(activity.ID)
            fmt.Printf("You removed the activity \"%s\"\n", name)
        }
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
