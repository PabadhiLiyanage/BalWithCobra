/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove the artifacts created during the build",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().String("target-dir", "", "Target directory path")
}
