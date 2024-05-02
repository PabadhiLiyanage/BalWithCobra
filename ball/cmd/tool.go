package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Extend the Ballerina CLI with custom commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tool called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(toolCmd)
	toolCmd.AddCommand(pullToolCmd)
	toolCmd.AddCommand(removeCmd)
	toolCmd.AddCommand(updateCmd)
	toolCmd.AddCommand(useCmd)
	toolCmd.AddCommand(listCmd)
	toolCmd.AddCommand(searchToolCmd)
	pullToolCmd.Flags().String("repository", "", "Pull a tool from local repository.")
	removeCmd.Flags().String("repository", "", " Remove a tool from a custom repository")
	useCmd.Flags().String("repository", "", " Use a tool from a custom repository")

}

var pullToolCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull a tool from the Ballerina Central",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
		executeBallerinaCommand()
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a tool from the local environment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		executeBallerinaCommand()
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a tool to the latest version compatible with the current distribution.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		executeBallerinaCommand()
	},
}

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set a tool version as the active version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use called")
		executeBallerinaCommand()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tools available in the local environment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		executeBallerinaCommand()
	},
}

var searchToolCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the Ballerina Central for tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		executeBallerinaCommand()
	},
}
