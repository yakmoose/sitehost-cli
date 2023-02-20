/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "dns",
	Short: "Commands for managing dns zones",
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
