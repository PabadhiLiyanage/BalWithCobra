/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// distCmd represents the dist command
var distCmd = &cobra.Command{
	Use:   "dist",
	Short: "Manage Ballerina distributions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("give control to update tool")
	},
}

func init() {
	RootCmd.AddCommand(distCmd)
	distCmd.Flags().BoolP("pre-releases", "p", false, "List the distributions that are available for you to download under the pre-releases channel")
	distCmd.Flags().BoolP("all", "a", false, "")
}
