/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domains",
	Short: "Create and delete domain names",
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
