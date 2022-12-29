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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all activities",
    Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
        var activities []models.Activity
        if err := db.Get().Preload("ActivePeriod").Find(&activities).Error; err != nil {
            log.Fatal(err)
        }
        for _, activity := range activities {
            activeStr := "\n"
            if activity.ActivePeriod != nil {
                activeStr = "(active)\n"
            }
            fmt.Printf("- %s %s", activity.Name, activeStr)
        }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
