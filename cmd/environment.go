/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var environmentCmd = &cobra.Command{
	Use:   "env",
	Short: "Commands for managing stack environment variables",
}

func init() {
	stacksCmd.AddCommand(environmentCmd)
	environmentCmd.PersistentFlags().StringP("stack", "s", "", "The project/stack to get")
	environmentCmd.MarkPersistentFlagRequired("stack")
	environmentCmd.PersistentFlags().StringP("service", "", "", "The service/container to get")
}
