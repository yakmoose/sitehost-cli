package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listStackServers = &cobra.Command{
	Use:   "list",
	Short: "List cloud servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := stackserver.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
		servers, err := client.List(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if format == "json" {
			json, err := json.MarshalIndent(servers, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Server Name\tServer Label")
			for _, stackServer := range *servers {
				fmt.Fprintf(w, "%s\t%s\n", stackServer.Name, stackServer.Label)
			}

			fmt.Fprintln(w)
		} else {
			/// how to error out.
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	stackServerCommand.AddCommand(listStackServers)
}
