/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ball/cmd"
	"fmt"
)

func main() {
	cmd.Execute()
	subCommands := make([]string, 0)

	// Append subcommands to the slice
	for _, cmd := range cmd.RootCmd.Commands() {
		subCommands = append(subCommands, cmd.Use)
	}

	// Print the subcommand list
	fmt.Println("Subcommands appended to the slice:")
	fmt.Println(subCommands)
}
