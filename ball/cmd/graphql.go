/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Generate the GraphQL schema for a Ballerina GraphQL service",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graphql called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(graphqlCmd)
	graphqlCmd.Flags().StringP("input", "i", "", "This is mandatory input. The given Ballerina GraphQL service file will generate the GraphQL schema (SDL) file relevant to the service")
	graphqlCmd.Flags().StringP("output", "o", "", "")
	graphqlCmd.Flags().StringP("service", "s", "", "")
	graphqlCmd.Flags().BoolP("use-records-for-objects", "r", false, "")
	graphqlCmd.Flags().BoolP("mode", "m", false, "")

}
