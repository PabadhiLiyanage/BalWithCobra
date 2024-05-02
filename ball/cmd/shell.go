package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Run Ballerina interactive REPL",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shell called")
	},
}

func init() {
	RootCmd.AddCommand(shellCmd)
	shellCmd.Flags().BoolP("debug", "d", false, "Enable debug mode from the beginning.")
	shellCmd.Flags().Bool("force-dumb", false, "Force the dumb terminal mode.")
	shellCmd.Flags().StringP("file", "f", "", "Open a file and load the initial declarations.")
	shellCmd.Flags().StringP("time-out", "t", "", "Set the tree parsing timeout value.")
}
