/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudCmd represents the cloud command
var stackServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Commands for manipulating with Sitehost cloud stack servers",
}

func init() {
	cloudCommand.AddCommand(stackServerCommand)
}
