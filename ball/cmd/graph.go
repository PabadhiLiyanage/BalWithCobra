/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Print the dependency graph",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graph called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(graphCmd)

	graphCmd.Flags().Bool("dump-raw-graphs", false, "Print all intermediate graphs created in the dependency resolution process.")
	graphCmd.Flags().Bool("offline", false, "Proceed without accessing the network.")
	graphCmd.Flags().Bool("sticky", false, "Attempt to stick to the dependency versions available in the 'Dependencies.toml' file.")
	graphCmd.Flags().String("target-dir", "", "Target directory path")
}
