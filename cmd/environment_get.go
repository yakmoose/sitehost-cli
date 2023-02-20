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
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack/environment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var environmentGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the stack environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := environment.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		serviceName := cmd.Flag("service").Value.String()

		format := cmd.Flag("format").Value.String()

		if "" == serviceName {
			serviceName = stackName
		}

		environmentVariables, err := client.Get(context.Background(), environment.GetRequest{ServerName: serverName, Project: stackName, Service: serviceName})
		if err != nil {
			return err
		}

		if format == "json" {
			json, err := json.MarshalIndent(environmentVariables, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Name\tValue")
			for _, variable := range *environmentVariables {
				fmt.Fprintf(w, "%s\t%s\t\n", variable.Name, variable.Content)
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
	environmentCmd.AddCommand(environmentGetCmd)
}
