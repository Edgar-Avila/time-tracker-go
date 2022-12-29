/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time-tracker/db"
	"time-tracker/models"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an activity",
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        if err := db.Get().Unscoped().Delete(&models.Activity{}, "name = ?", name).Error; err != nil {
            log.Fatal(err)
        }
        fmt.Printf("You removed the activity \"%s\"", name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
