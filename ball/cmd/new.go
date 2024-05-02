/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
		executeBallerinaCommand()
	},
}

func init() {

	RootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("template", "t", "", "Create a package using a predefined template.")

}
