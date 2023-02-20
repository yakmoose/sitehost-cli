/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Commands for managing dns zone records",
}

func init() {
	domainCmd.AddCommand(recordCmd)
	recordCmd.PersistentFlags().StringP("domain", "d", "", "The domain zone to use")
	recordCmd.MarkFlagRequired("domain")
}
