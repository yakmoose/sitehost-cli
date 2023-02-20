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
var startStackCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a running stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("apiKey")
		clientId := viper.GetString("clientId")
		client := stack.New(api.NewClient(apiKey, clientId))

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()
		job, err := client.Start(context.Background(), stack.StopStartRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		return helper.WaitForAction(api.NewClient(apiKey, clientId), job.Return.JobID)
	},
}

func init() {
	stackCommand.AddCommand(startStackCmd)

	startStackCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	startStackCmd.MarkFlagRequired("server")

	startStackCmd.Flags().StringP("stack", "s", "", "The stack to start")
	startStackCmd.MarkFlagRequired("stack")
}
