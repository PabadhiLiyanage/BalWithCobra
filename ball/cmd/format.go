/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format the Ballerina source files",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("format called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(formatCmd)
	formatCmd.Flags().BoolP("dry-run", "d", false, "Perform a dry run of the formatter and see which files will be formatted after the execution")
}
