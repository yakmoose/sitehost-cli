/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shcli/pkg/cloud/stacks"
)

// listCmd represents the list command
var listStacksCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud stacks on a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := stacks.StackClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		serverName := cmd.Flag("server").Value.String()
		stacks, err := client.List(context.Background(), &stack.ListRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(stacks, "", "  ")

		fmt.Println(string(json))

		return nil
	},
}

func init() {
	stacksCmd.AddCommand(listStacksCmd)
}
