/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stackCommand represents the stacks command
var stackCommand = &cobra.Command{
	Use:   "stack",
	Short: "Manage cloud stacks",
}

func init() {
	cloudCommand.AddCommand(stackCommand)
}
