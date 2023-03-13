/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/db"
	"github.com/sitehostnz/terraform-provider-sitehost/sitehost/helper"
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

		dbDeleteResponse, err := dbClient.Delete(ctx, db.DeleteRequest{Database: database, MySQLHost: host, ServerName: serverName})
		if err != nil {
			return err
		}

		return helper.WaitForAction(client, dbDeleteResponse.Return.JobID)
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
