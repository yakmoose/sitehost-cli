/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"errors"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/sitehostnz/gosh/pkg/api/server"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudDbAdd represents the domainAdd command
var cloudDbAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a new database",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		client := api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId"))
		dbClient := db.New(client)
		serverClient := server.New(client)

		database := cmd.Flag("db").Value.String()
		serverName := cmd.Flag("server").Value.String()
		host := cmd.Flag("host").Value.String()
		container := cmd.Flag("container").Value.String()

		// 1. is the server a stack server? best check
		stackServer, err := serverClient.Get(context.Background(), server.GetRequest{ServerName: serverName})
		if err != nil {
			return err
		}

		if stackServer.Server.ProductType != "CLDCON" {
			return errors.New("server is not a cloud container server")
		}

		dbAddResponse, err := dbClient.Add(ctx, db.AddRequest{Database: database, MySQLHost: host, ServerName: serverName, Container: container})
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		return helper.WaitForAction(client, dbAddResponse.Return.JobID)
	},
}

func init() {
	cloudDbCmd.AddCommand(cloudDbAdd)

	cloudDbAdd.Flags().StringP("server", "S", "", "The server name")
	cloudDbAdd.MarkFlagRequired("server")

	cloudDbAdd.Flags().StringP("host", "H", "", "The database host")
	cloudDbAdd.MarkFlagRequired("host")

	cloudDbAdd.Flags().StringP("db", "d", "", "The database name")
	cloudDbAdd.MarkFlagRequired("db")

	cloudDbAdd.Flags().StringP("container", "c", "", "The database backup container")
	cloudDbAdd.MarkFlagRequired("container")

}
