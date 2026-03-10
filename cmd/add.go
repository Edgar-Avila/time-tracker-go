package cmd

import (
	"time-tracker/models"
	"time-tracker/repo"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an activty",
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			repo.ActivityRepo().Create(&models.Activity{Name: name})
			color.New(color.FgGreen).Printf("You added a new activity \"%s\"\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
