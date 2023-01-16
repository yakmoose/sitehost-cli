/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shcli/pkg/domains"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all domain names",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := domains.DomainClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		domains, err := client.List(context.Background())
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(domains, "", "  ")

		fmt.Println(string(json))

		return nil
	},
}

func init() {
	domainCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
