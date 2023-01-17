/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shcli/pkg/cloud/stacks"
)

// listCmd represents the list command
var restartStackCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("apiKey")
		clientId := viper.GetString("clientId")
		client := stacks.StackClient(apiKey, clientId)
		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		job, err := client.Restart(context.Background(), &stack.StopStartRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		return helper.WaitForAction(api.NewClient(apiKey, clientId), job.Return.JobID)
	},
}

func init() {
	stacksCmd.AddCommand(restartStackCmd)
	restartStackCmd.Flags().StringP("stack", "s", "", "The server name to fetch the stack list from")
	restartStackCmd.MarkFlagRequired("stack")
}
