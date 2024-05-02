/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Ballerina version",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
