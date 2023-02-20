package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/dns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listRecordsCmd = &cobra.Command{
	Use:   "list",
	Short: "List zones records",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := dns.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		domainName := cmd.Flag("domain").Value.String()

		records, err := client.ListRecords(context.Background(), dns.ListRecordsRequest{Domain: domainName})
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if format == "json" {
			json, err := json.MarshalIndent(records, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Id\tDomain\tName\tType\tValue\tPriority")
			for _, record := range records.Return {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", record.ID, record.Domain, record.Name, record.Type, record.Content, record.Priority)
			}

			fmt.Fprintln(w)
		} else {
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	recordCmd.AddCommand(listRecordsCmd)
}
