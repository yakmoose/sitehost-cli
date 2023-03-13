/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbUpdate updates the cloud db backup container
var cloudDbUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update the database backup location",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		client := db.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		database := cmd.Flag("db").Value.String()
		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()
		container := cmd.Flag("container").Value.String()

		// don't check on update, as if we change this, will assume a not found?

		response, err := client.Update(
			ctx,
			db.UpdateRequest{
				Database:   database,
				MySQLHost:  host,
				ServerName: serverName,
				Container:  container,
			})
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	cloudDbCmd.AddCommand(cloudDbUpdate)

	cloudDbUpdate.Flags().StringP("server", "S", "", "The server name")
	cloudDbUpdate.MarkFlagRequired("server")

	cloudDbUpdate.Flags().StringP("host", "H", "", "The database host")
	cloudDbUpdate.MarkFlagRequired("host")

	cloudDbUpdate.Flags().StringP("db", "d", "", "The database name")
	cloudDbUpdate.MarkFlagRequired("db")

	cloudDbUpdate.Flags().StringP("container", "c", "", "The database backup container")
	cloudDbUpdate.MarkFlagRequired("container")

}
