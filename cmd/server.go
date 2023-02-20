/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cloudCmd represents the cloud command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Commands for manipulating with Sitehost servers",
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
