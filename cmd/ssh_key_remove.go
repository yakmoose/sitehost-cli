/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/ssh/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command.
var removeKeyCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove the specified ssh key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := key.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		keyID := cmd.Flag("keyid").Value.String()

		keyResponse, err := client.Remove(context.Background(), key.RemoveRequest{ID: keyID})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(keyResponse, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	sshKeysCmd.AddCommand(removeKeyCmd)

	removeKeyCmd.Flags().StringP("keyid", "", "", "The key id")
	removeKeyCmd.MarkFlagRequired("keyid")
}
