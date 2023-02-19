/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import "github.com/spf13/cobra"

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage ssh",
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
