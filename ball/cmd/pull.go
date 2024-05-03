/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch packages from Ballerina Central or a custom package repository",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
		executeBallerinaCommand()

	},
}

func init() {
	RootCmd.AddCommand(pullCmd)
	pullCmd.Flags().String("repository", "", " Pull a package from a custom repository.")
}
