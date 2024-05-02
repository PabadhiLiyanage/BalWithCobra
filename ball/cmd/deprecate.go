package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deprecateCmd = &cobra.Command{
	Use:   "deprecate",
	Short: "Deprecates a published package",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deprecate called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(deprecateCmd)
	deprecateCmd.Flags().String("message", "", "Use the given <msg> as the deprecation message")
	deprecateCmd.Flags().Bool("undo", false, "Undo deprecation of a package")
}
