/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/spf13/viper"
	"shcli/pkg/domains"

	"github.com/sitehostnz/gosh/pkg/api/domain"
	"github.com/spf13/cobra"
)

// domainAddCmd represents the domainAdd command
var domainDelCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a domain name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := domains.DomainClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		ctx := context.Background()
		domain, err := client.Get(ctx, domain.GetRequest{DomainName: args[0]})
		if err != nil {
			return err
		}

		if domain == nil {
			return nil
		}

		domain, err = client.Delete(ctx, &models.Domain{Name: args[0], TemplateID: "0"})
		if err != nil {
			return err
		}

		// json, err := json.MarshalIndent(domain, "", "  ")
		// fmt.Println()

		return nil
	},
}

func init() {
	domainCmd.AddCommand(domainDelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
