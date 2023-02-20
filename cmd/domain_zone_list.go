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
	"os"
	"text/tabwriter"
)

// listZonesCommand represents the list command
var listZonesCommand = &cobra.Command{
	Use:   "list",
	Short: "List zones records",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		zones, err := client.ListZones(context.Background(), &dns.ListZoneOptions{})
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if format == "json" {
			json, err := json.MarshalIndent(zones, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Domain")
			for _, zone := range *zones {
				fmt.Fprintln(w, zone.Name)
			}

			fmt.Fprintln(w)
		} else {
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	domainCmd.AddCommand(listZonesCommand)
}
