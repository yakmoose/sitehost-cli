/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudImageCommand represents the stacks command
var cloudDbCmd = &cobra.Command{
	Use:   "db",
	Short: "Manage cloud databases",
}

func init() {
	cloudCommand.AddCommand(cloudDbCmd)
}
