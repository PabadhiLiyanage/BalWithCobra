/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var semverCmd = &cobra.Command{
	Use:   "semver",
	Short: "Validate SemVer compliance of the package changes",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("semver called")
		executeBallerinaCommand()
	},
}

func init() {
	RootCmd.AddCommand(semverCmd)
	semverCmd.Flags().BoolP("show-diff", "d", false, "Display the detailed information of source code differences between the compared package versions")
	semverCmd.Flags().StringP("compare-version", "c", "", "Analyze SemVer compliance of the local changes against the specified release version")
}
