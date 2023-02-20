/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/spf13/viper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
)

// domainAddCmd represents the domainAdd command
var domainDelCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a domain name",
	RunE: func(cmd *cobra.Command, args []string) error {
		domainName := cmd.Flag("domain").Value.String()
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()
		domainGetResponse, err := client.GetZone(ctx, dns.GetZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if domainGetResponse.Return == nil {
			return nil
		}

		_, err = client.DeleteZone(ctx, dns.DeleteZoneRequest{DomainName: domainName})
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
