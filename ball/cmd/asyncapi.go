/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// asyncapiCmd represents the asyncapi command
var asyncapiCmd = &cobra.Command{
	Use:   "asyncapi",
	Short: "Generate a Ballerina listener from an AsyncAPI contract.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asyncapi called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(asyncapiCmd)
	asyncapiCmd.Flags().StringP("input", "i", "", "This is a mandatory input. The listener will be generated according to the given AsyncAPI contract.")
	asyncapiCmd.Flags().StringP("output", "o", "", "Location of the generated Ballerina source code.")
}
