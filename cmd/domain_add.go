/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/spf13/viper"
	"shcli/pkg/domains"

	"github.com/sitehostnz/gosh/pkg/api/domain"
	"github.com/spf13/cobra"
)

// domainAddCmd represents the domainAdd command
var domainAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := domains.DomainClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		ctx := context.Background()

		domainName := cmd.Flag("domain").Value.String()

		domain, err := client.Get(ctx, domain.GetRequest{DomainName: domainName})
		if err != nil {
			return err
		}

		if domain != nil {
			return errors.New("Domain already exists")
		}

		domain, err = client.Create(ctx, &models.Domain{Name: domainName, TemplateID: "0"})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(domain, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	domainCmd.AddCommand(domainAddCmd)
	domainAddCmd.Flags().StringP("domain", "d", "", "The domain name to use")
	domainAddCmd.MarkFlagRequired("domain")

}
