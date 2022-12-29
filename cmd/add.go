/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an activty",
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        repo.ActivityRepo().Create(&models.Activity{ Name: name, })
        fmt.Printf("You added a new activity \"%s\"", name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
