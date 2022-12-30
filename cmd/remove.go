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
	Run: func(cmd *cobra.Command, args []string) {
        for _, name := range args {
            repo.ActivityRepo().DeleteByName(name)
            fmt.Printf("You removed the activity \"%s\"\n", name)
        }
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
