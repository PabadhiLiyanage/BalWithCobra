package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new module to the current package",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

}
