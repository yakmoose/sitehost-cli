/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// domainAddCmd represents the domainAdd command
var domainAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		ctx := context.Background()

		domainName := cmd.Flag("domain").Value.String()

		domainResponse, err := client.GetZone(ctx, dns.GetZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if domainResponse.Return != nil {
			return errors.New("Domain already exists")
		}

		zoneCreateResponse, err := client.CreateZone(ctx, dns.CreateZoneRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(zoneCreateResponse.Return, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	domainCmd.AddCommand(domainAddCmd)
	domainAddCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	domainAddCmd.MarkFlagRequired("domain")

}
