/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var serverGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the specified server",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := server.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()

		stack, err := client.Get(context.Background(), server.GetRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	serverCmd.AddCommand(serverGetCmd)
	serverGetCmd.Flags().StringP("server", "S", "", "The server name to get")
	serverGetCmd.MarkPersistentFlagRequired("server")
}
