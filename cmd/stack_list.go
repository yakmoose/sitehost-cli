/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listStacksCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud stacks on a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := stack.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()
		stacks, err := client.List(context.Background(), stack.ListRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()

		if format == "json" {
			json, err := json.MarshalIndent(stacks, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Stack Name\tStack Label\tServer")
			for _, stack := range *stacks {
				fmt.Fprintf(w, "%s\t%s\t%s\n", stack.Name, stack.Label, stack.Server)
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
	stackCommand.AddCommand(listStacksCmd)
	listStacksCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	listStacksCmd.MarkFlagRequired("server")

}
