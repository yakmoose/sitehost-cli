/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var stopStackCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a running stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("apiKey")
		clientId := viper.GetString("clientId")
		client := stack.New(api.NewClient(apiKey, clientId))

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		job, err := client.Start(context.Background(), stack.StopStartRestartRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		return helper.WaitForAction(api.NewClient(apiKey, clientId), job.Return.JobID)
	},
}

func init() {
	stackCommand.AddCommand(stopStackCmd)

	stopStackCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	stopStackCmd.MarkFlagRequired("server")

	stopStackCmd.Flags().StringP("stack", "s", "", "The stack to stop")
	stopStackCmd.MarkFlagRequired("stack")
}
