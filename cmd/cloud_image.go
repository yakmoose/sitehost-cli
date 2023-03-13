/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudImageCommand represents the stacks command
var cloudImageCommand = &cobra.Command{
	Use:   "image",
	Short: "Manage cloud images",
}

func init() {
	cloudCommand.AddCommand(cloudImageCommand)
}
