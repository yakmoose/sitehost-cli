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
	"github.com/sitehostnz/gosh/pkg/api/cloud/server"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

type VhostList struct {
	VHOST, Name, Label, Server, ServerLabel string
}

// listStackVirtualHostsCmd list all virtual hosts
var listStackVirtualHostsCmd = &cobra.Command{
	Use:   "virtualhosts",
	Short: "List virtual stacks on a server",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		stackServerClient := server.New(client)
		stackClient := stack.New(client)

		stackServersResponse, err := stackServerClient.List(context.Background())
		if err != nil {
			return err
		}

		pattern := "'VIRTUAL_HOST=(.*)'"
		expr, _ := regexp.Compile(pattern)

		vhosts := []VhostList{}
		for _, ss := range stackServersResponse.CloudServers {
			stacksResponse, err := stackClient.List(context.Background(), stack.ListRequest{ServerName: ss.Name})
			if err != nil {
				return err
			}

			for _, stack := range stacksResponse.Return.Stacks {
				m := expr.FindStringSubmatch(stack.DockerFile)
				if len(m) == 2 {
					for _, v := range strings.Split(m[1], ",") {
						vhosts = append(vhosts, VhostList{VHOST: v, Name: stack.Name, Label: stack.Label, Server: stack.Server, ServerLabel: stack.ServerLabel})
					}
				}
			}
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(vhosts, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Vhost\tStack Name\tStack Label\tServer Name\tServer Label")
			for _, stack := range vhosts {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", stack.VHOST, stack.Name, stack.Label, stack.Server, stack.ServerLabel)
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
	stackCommand.AddCommand(listStackVirtualHostsCmd)
}
