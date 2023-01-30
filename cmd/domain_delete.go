/*
Copyright © 2023 John Lennard <john@yakmoo.se>
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
	RunE: func(cmd *cobra.Command, args []string) error {
		domainName := cmd.Flag("domain").Value.String()
		client := domains.DomainClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		ctx := context.Background()
		domain, err := client.Get(ctx, domain.GetRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if domain == nil {
			return nil
		}

		domain, err = client.Delete(ctx, &models.Domain{Name: domainName, TemplateID: "0"})
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	domainCmd.AddCommand(domainDelCmd)
	domainDelCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	domainDelCmd.MarkFlagRequired("domain")
}
