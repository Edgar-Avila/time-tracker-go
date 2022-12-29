/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        repo.ActivityRepo().DeleteByName(name)
        fmt.Printf("You removed the activity \"%s\"", name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
