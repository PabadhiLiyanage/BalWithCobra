/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the Ballerina Archive (BALA) of the current package to a package repository",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(pushCmd)
	pushCmd.Flags().String("repository", "", " Push the BALA of the current package to a custom repository.")
}
