/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shcli/pkg/cloud/stacks"
)

// listCmd represents the list command
var startStackCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a running stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := stacks.StackClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		_, err := client.Start(context.Background(), &stack.StopStartRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}
		
		return nil
	},
}

func init() {
	stacksCmd.AddCommand(startStackCmd)
	startStackCmd.Flags().StringP("stack", "s", "", "The stack to start")
	startStackCmd.MarkFlagRequired("stack")
}
