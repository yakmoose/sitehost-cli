/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudImageCommand represents the stacks command
var stackImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage stack images",
}

func init() {
	stackCommand.AddCommand(stackImageCmd)
}
