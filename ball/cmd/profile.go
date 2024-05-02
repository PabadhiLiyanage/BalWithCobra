/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("profile called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(profileCmd)
	profileCmd.Flags().String("debug", "", "Run Ballerina Profiler in the remote debugging mode.")
}
