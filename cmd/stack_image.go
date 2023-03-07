/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stackCommand represents the stacks command
var stackImageCommand = &cobra.Command{
	Use:   "image",
	Short: "Manage cloud images",
}

func init() {
	cloudCommand.AddCommand(stackImageCommand)
}
