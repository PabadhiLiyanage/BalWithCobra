/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
	testCmd.Flags().Bool("code-coverage", false, "give code coverage in tests")
	testCmd.Flags().String("coverage-format", "", "Generate a coverage report in the specified format.")
	testCmd.Flags().Int("debug", 0, "Run in the remote debugging mode")
	testCmd.Flags().String("disable-groups", "", "Specify the test groups to be excluded")
	testCmd.Flags().String("groups", "", " Specify the test groups to be executed.")
	testCmd.Flags().Bool("list-groups", false, "List the test groups available in the test files")
	testCmd.Flags().Bool("offline", false, "Proceed without accessing the network")
	testCmd.Flags().Bool("observability-included", false, "Include the dependencies that are required to enable observability")
	testCmd.Flags().Bool("sticky", false, "Attempt to stick to the dependency versions available in the'Dependencies.toml' file")
	testCmd.Flags().Bool("rerun-failed", false, "give code coverage in tests")
	testCmd.Flags().String("target-dir", "", "specify the target directory")
	testCmd.Flags().Bool("test-report", false, "Generate an HTML report containing the test results")
	testCmd.Flags().String("tests", "", "Specify the test functions to be executed and Specify the test functions to be executed only from the given module.")
	testCmd.Flags().Bool("graalvm", false, "Execute test cases using the GraalVM native image")
	testCmd.Flags().String("graalvm-build-options", "", "Additional build options to be passed to the GraalVM native image")
	testCmd.Flags().String("excludes", "", "Exclude a specific Ballerina source files or folders from code coverage calculation")

}
