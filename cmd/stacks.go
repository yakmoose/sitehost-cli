/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stacksCmd represents the stacks command
var stacksCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage cloud stacks",
}

func init() {
	cloudCmd.AddCommand(stacksCmd)
}
