/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doc called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(docCmd)
	docCmd.Flags().StringP("output", "o", "", " Write the output to the given file.")
	docCmd.Flags().StringP("exclude", "e", "", " Write the output to the given file.")
	docCmd.Flags().Bool("offline", false, "Run in offline mode")
	docCmd.Flags().Bool("sticky", false, "Run with including sticky")
	docCmd.Flags().String("target-dir", "", "specify the target directory")

}
