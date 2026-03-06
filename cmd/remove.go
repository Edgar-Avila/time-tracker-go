package cmd

import (
	"github.com/spf13/cobra"
)

// removeCmd is the parent for remove subcommands
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove activities or records",
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
