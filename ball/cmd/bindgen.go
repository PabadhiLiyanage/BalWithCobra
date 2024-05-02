package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bindgenCmd = &cobra.Command{
	Use:   "bindgen",
	Short: "Generate Ballerina bindings for Java APIs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bindgen called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(bindgenCmd)
	bindgenCmd.Flags().String("classpath", "", "")
	bindgenCmd.Flags().String("maven", "", "A Maven dependency with a colon-delimited group ID, artifact ID, and version.")
	bindgenCmd.Flags().StringP("output", "o", "", "")
	bindgenCmd.Flags().Bool("public", false, "")
	bindgenCmd.Flags().Bool("with-optional-types", false, "")
	bindgenCmd.Flags().Bool("with-optional-types-param", false, "")
	bindgenCmd.Flags().Bool("with-optional-types-return", false, "")
}
