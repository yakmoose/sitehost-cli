/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/info"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Display sitehost API info",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := info.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))
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
