/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/ssh/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"strings"

	"context"
	"os"
)

// updateKeyCmd updates the spcified key with the supplied details
var updateKeyCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a ssh key",
	RunE: func(cmd *cobra.Command, args []string) error {

		label := cmd.Flag("label").Value.String()
		keyFileName := cmd.Flag("key").Value.String()
		keyID := cmd.Flag("keyid").Value.String()

		// read the docker file.
		var fd *os.File
		var err error
		if len(keyFileName) > 0 {
			fd, err = os.Open(keyFileName)
			if nil != err {
				return err
			}
		} else {
			fd = os.Stdin
		}

		keyFile, err := io.ReadAll(fd)

		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		keyClient := key.New(client)

		// 1. is the server a stack server? best check
		keyResponse, err := keyClient.Update(context.Background(), key.UpdateRequest{ID: keyID, Content: strings.TrimSpace(string(keyFile)), Label: label})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(keyResponse, "", "  ")
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	sshKeysCmd.AddCommand(updateKeyCmd)

	updateKeyCmd.Flags().StringP("label", "l", "", "The key label")
	updateKeyCmd.MarkFlagRequired("label")

	updateKeyCmd.Flags().StringP("keyid", "", "", "The key id")
	updateKeyCmd.MarkFlagRequired("keyid")

	updateKeyCmd.Flags().StringP("key", "k", "", "The key file")

}
