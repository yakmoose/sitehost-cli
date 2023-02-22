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
	"regexp"
	"strings"
	"text/tabwriter"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/server"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/gosh/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func formatStackOutput(format string, stack models.Stack) error {

	if format == "json" {
		json, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))
	} else if format == "text" {

		pattern := "'VIRTUAL_HOST=(.*)'"

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 4, 4, ' ', 1)
		fmt.Fprintln(w, "Label\tValue")
		fmt.Fprintf(w, "Label\t%s\n", stack.Label)
		fmt.Fprintf(w, "Name\t%s\n", stack.Name)
		fmt.Fprintf(w, "Server Name\t%s\n", stack.Server)
		fmt.Fprintf(w, "Server Label\t%s\n", stack.ServerLabel)
		expr, _ := regexp.Compile(pattern)
		m := expr.FindStringSubmatch(stack.DockerFile)
		if len(m) == 2 {
			fmt.Fprintf(w, "Aliases\t%s\n", strings.Join(strings.Split(m[1], ","), ", "))
		}
		fmt.Fprintf(w, "Admin URL\thttps://cp.sitehost.nz/cloud/manage-container/server/%s/stack/%s\n", stack.Server, stack.Name)
		fmt.Fprintln(w)

	} else {
		/// how to error out.
		return errors.New("unsupported format, please choose text or json")
	}

	return nil
}

// listCmd represents the list command
var findStack = &cobra.Command{
	Use:   "find",
	Short: "Find a stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		stackServerClient := stackserver.New(client)
		stackClient := stack.New(client)

		stackServers, err := stackServerClient.List(context.Background())
		if err != nil {
			return err
		}

		pattern := "'VIRTUAL_HOST=.*(" + regexp.QuoteMeta(args[0]) + ").*'"
		expr, _ := regexp.Compile(pattern)

		for _, ss := range *stackServers {
			stacks, err := stackClient.List(context.Background(), stack.ListRequest{ServerName: ss.Name})
			if err != nil {
				return err
			}

			for _, s := range *stacks {
				// check the name,
				if s.Name == args[0] || s.Label == args[0] {
					return formatStackOutput(cmd.Flag("format").Value.String(), s)
				}

				// then check the aliases/vhost stuff
				// get the virual host and split
				m := expr.Match([]byte(s.DockerFile))
				if m {
					return formatStackOutput(cmd.Flag("format").Value.String(), s)
				}

			}

		}
		return nil
	},
}

func init() {
	stackCommand.AddCommand(findStack)
}
