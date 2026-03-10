package cmd

import (
	"strconv"
	"time-tracker/repo"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// remove record by id
var removeRecordCmd = &cobra.Command{
	Use:   "record [id]",
	Short: "Remove a record by id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			color.New(color.FgRed).Printf("invalid id: %v\n", err)
			return
		}

		if err := repo.RecordRepo().DeleteByID(uint(id)); err != nil {
			color.New(color.FgRed).Printf("failed to remove record: %v\n", err)
			return
		}

		color.New(color.FgGreen).Printf("Removed record %d\n", id)
	},
}

func init() {
	removeCmd.AddCommand(removeRecordCmd)
}
