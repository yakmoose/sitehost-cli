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

// cloudDbDelete represents the domainAdd command
var cloudDbDelete = &cobra.Command{
	Use:   "rm",
	Short: "delete a database",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		dbClient := db.New(client)

		database := cmd.Flag("db").Value.String()
		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()
		container := cmd.Flag("container").Value.String()

		dbDeleteResponse, err := dbClient.Add(ctx, db.AddRequest{Database: database, MySQLHost: host, ServerName: serverName, Container: container})
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(dbDeleteResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	cloudDbCmd.AddCommand(cloudDbDelete)

	cloudDbDelete.Flags().StringP("server", "S", "", "The server name")
	cloudDbDelete.MarkFlagRequired("server")

	cloudDbDelete.Flags().StringP("host", "H", "", "The database host")
	cloudDbDelete.MarkFlagRequired("host")

	cloudDbDelete.Flags().StringP("db", "d", "", "The database name")
	cloudDbDelete.MarkFlagRequired("db")

}
