/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Ballerina Central for packages",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

}
