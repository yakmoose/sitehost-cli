/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/stack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stackGetDockerFileCmd represents the get command
var stackGetDockerFileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Get the dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := stack.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		serverName := cmd.Flag("server").Value.String()
		stackName := cmd.Flag("stack").Value.String()

		stack, err := client.Get(context.Background(), stack.GetRequest{ServerName: serverName, Name: stackName})
		if err != nil {
			return err
		}

		fmt.Println(stack.DockerFile)

		return nil
	},
}

func init() {
	stackCommand.AddCommand(stackGetDockerFileCmd)

	stackGetDockerFileCmd.Flags().StringP("server", "S", "", "The server name to fetch the stack list from")
	stackGetDockerFileCmd.MarkFlagRequired("server")

	stackGetDockerFileCmd.Flags().StringP("stack", "s", "", "The stack to get")
	stackGetDockerFileCmd.MarkFlagRequired("stack")
}
