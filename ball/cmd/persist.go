package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var persistCmd = &cobra.Command{
	Use:   "persist",
	Short: "Manage data persistence",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("persist called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(persistCmd)
	persistCmd.AddCommand(initCmd)
	persistCmd.AddCommand(generateCmd)
	initCmd.Flags().String("module", "", "Initialize the Ballerina package for persistence with the given module")
	initCmd.Flags().String("datastore", "", "The type of the datastore to be used for persistence")

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Ballerina package for persistence.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		executeBallerinaCommand()
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the client API based on the data model.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		executeBallerinaCommand()
	},
}
