/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import "github.com/spf13/cobra"

// stackCommand represents the stacks command
var sshKeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage ssh keys",
}

func init() {
	sshCmd.AddCommand(sshKeysCmd)
}
