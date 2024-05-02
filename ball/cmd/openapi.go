/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// openapiCmd represents the openapi command
var openapiCmd = &cobra.Command{
	Use:   "openapi",
	Short: "Generate a Ballerina service or a client from an OpenAPI contract and vice versa.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("openapi called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(openapiCmd)
	openapiCmd.Flags().StringP("input", "i", "", "This is mandatory input. The given Ballerina GraphQL service file will generate the GraphQL schema (SDL) file relevant to the service")
	openapiCmd.Flags().StringP("output", "o", "", "Location of the generated Ballerina source code.")
	openapiCmd.Flags().String("mode", "", "The Ballerina service and client will be generated according to the specified mode.")
	openapiCmd.Flags().BoolP("nullable", "n", false, "This is a safe option to generate all data types in the record with Ballerina nil support")
	openapiCmd.Flags().String("license", "", "Optional. The `.bal` files will generate with the given copyright or license header.")
	openapiCmd.Flags().StringP("service", "s", "", "This option is used with the command of Ballerina to OpenAPI specification generation.")
	openapiCmd.Flags().Bool("json", false, "Generate the Ballerina service to OpenAPI output in JSON.The default is YAML.")
	openapiCmd.Flags().String("tags", "", "This option is used with the OpenAPI to Ballerina file generation command.")
	openapiCmd.Flags().String("operations", "", "List of operations to generate the Ballerina service or client.")
	openapiCmd.Flags().Bool("with-tests", false, "Work with the client generation command and generate a test")
	openapiCmd.Flags().Bool("without-data-binding", false, "This option can be used in the service generation to generate a low-level service without any data-binding logic.")
	openapiCmd.Flags().String("client-methods", "", "This option can be used in client generation to select the client method type, which can be `resource` or `remote`.")

}
