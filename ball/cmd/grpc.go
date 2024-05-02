/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Generate Ballerina sources for the given Protocol Buffer definition",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().String("input", "", "Path to a '.proto' file or a directory containing multiple '.proto' files.")
	grpcCmd.Flags().String("output", "", "Location of the generated Ballerina source files.")
	graphCmd.Flags().String("proto-path", "", "Path to a directory in which to look for '.proto' files when resolving import directives.")
	graphCmd.Flags().Bool("mode", false, "Set the 'client' or 'service' mode to generate sample code.")
}
