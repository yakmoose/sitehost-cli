/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudCmd represents the cloud command
var cloudCommand = &cobra.Command{
	Use:   "cloud",
	Short: "Commands for manipulating with Sitehost cloud/container stacks",
}

func init() {
	rootCmd.AddCommand(cloudCommand)
}
