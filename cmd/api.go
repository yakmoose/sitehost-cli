/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shcli/pkg/api_info"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Display sitehost API details",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api_info.ApiClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		apiInfo, err := client.Get(context.Background())
		json, err := json.MarshalIndent(apiInfo, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
