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

// getCmd represents the get command
var stackGenerateName = &cobra.Command{
	Use:   "generatename",
	Short: "Generate a stack name",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := stack.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		stack, err := client.GenerateName(context.Background())
		if err != nil {
			return err
		}

		fmt.Println(stack.Return.Name)

		return nil
	},
}

func init() {
	stackCommand.AddCommand(stackGenerateName)
}
