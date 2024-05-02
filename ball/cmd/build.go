package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "ballerina-build - Compiles the current package",
	Example: "Build the current package. This will generate an 'app.jar' file in the'target/bin' directory.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
		executeBallerinaCommand()

	},
}

func init() {

	RootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringP("output", "o", "", " Write the output to the given file.")
	buildCmd.Flags().Bool("offline", false, "Run in offline mode")
	buildCmd.Flags().Bool("observability-included", false, "Run with including observability")
	buildCmd.Flags().Bool("sticky", false, "Run with including sticky")
	buildCmd.Flags().String("target-dir", "", "specify the target directory")
	buildCmd.Flags().Bool("export-openapi", false, "Run in offline mode")
	buildCmd.Flags().Bool("list-conflicted-classes", false, "Run in offline mode")
	buildCmd.Flags().Bool("graalvm", false, "Run in offline mode")
	buildCmd.Flags().String("graalvm-build-options", "", "Run in offline mode")
	buildCmd.Flags().String("cloud", "", "specify the target directory")

}
