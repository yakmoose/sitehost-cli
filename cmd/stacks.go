/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stacksCmd represents the stacks command
var stacksCmd = &cobra.Command{
	Use:   "stacks",
	Short: "Manage cloud stacks",
}

func init() {
	cloudCmd.AddCommand(stacksCmd)
}
