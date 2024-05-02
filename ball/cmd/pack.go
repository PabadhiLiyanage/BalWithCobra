/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pack called")
		executeBallerinaCommand()
	},
}

func init() {

	RootCmd.AddCommand(packCmd)
	packCmd.Flags().Bool("offline", false, "Run in offline mode")
	packCmd.Flags().Bool("sticky", false, "Run with including sticky")
	packCmd.Flags().String("target-dir", "", "specify the target directory")
}
