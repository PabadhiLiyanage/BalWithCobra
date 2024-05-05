package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "short description of health",
	Long:  "long description of health",
	Run: func(cmd *cobra.Command, args []string) {
		executeBallerinaCommand()
		fmt.Println("running")
	},
}

func init() {
	RootCmd.AddCommand(healthCmd)
	//Add subcommands
	healthCmd.AddCommand(sub1healthCmd)
	healthCmd.AddCommand(sub2healthCmd)

	RootCmd.Flags().StringP("output", "o", "", "Write the output to the given file.")
	RootCmd.Flags().StringP("mode", "m", "", "Mode can be 'package' or 'template'")

	// Add subcommand flags here
	sub1healthCmd.Flags().String("repository", "", "Use a tool from a custom repository.")

}

var sub1healthCmd = &cobra.Command{
	Use:   "sub1",
	Short: "this is short description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		executeBallerinaCommand()
	},
}

var sub2healthCmd = &cobra.Command{
	Use:   "sub2",
	Short: "this is short description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		executeBallerinaCommand()
	},
}
