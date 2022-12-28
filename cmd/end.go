/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "End an activity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("end called")
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
}
