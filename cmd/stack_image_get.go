/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/image"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var stackImageGetCommand = &cobra.Command{
	Use:   "get",
	Short: "Get the stack",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := image.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		imageCode := cmd.Flag("code").Value.String()

		stack, err := client.Get(context.Background(), image.GetRequest{Code: imageCode})
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
	stackImageCommand.AddCommand(stackImageGetCommand)

	stackImageGetCommand.Flags().StringP("code", "c", "", "The image code")
	stackImageGetCommand.MarkFlagRequired("code")

}
