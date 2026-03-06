package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd is the parent for list subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List activities or records",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
