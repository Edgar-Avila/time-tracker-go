/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time-tracker/repo"
	"time-tracker/util"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "View a report of your activities",
	Run: func(cmd *cobra.Command, args []string) {
        // Get flags
        timespan, err := cmd.Flags().GetString("timespan")
        if err != nil {
            log.Fatal(err)
        }

        // Check if timespan is valid
        validValues := []string{"all", "day", "week", "month", "year"}
        isValid := util.StringInSlice(timespan, validValues)
        if !isValid {
            fmt.Printf("Timespan should be one of %v", validValues)
        }

        // Get results
        results := repo.PeriodRepo().GetAfter(timespan)
        for _, result := range results {
            util.PrettyPrint(result)
        }
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().StringP("timespan", "t", "all", "How old should the analytics be?")
}
