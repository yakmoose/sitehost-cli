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

// cloudImageGetCommand represents the get command
var cloudDbGetCommand = &cobra.Command{
	Use:   "get",
	Short: "Get the database",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := db.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		databaseResponse, err := client.Get(
			context.Background(),
			db.GetRequest{
				Database:   cmd.Flag("db").Value.String(),
				ServerName: cmd.Flag("server").Value.String(),
				MySQLHost:  cmd.Flag("host").Value.String(),
			})

		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(databaseResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(json))

		return nil
	},
}

func init() {
	cloudDbCmd.AddCommand(cloudDbGetCommand)

	cloudDbGetCommand.Flags().StringP("server", "S", "", "The server name")
	cloudDbGetCommand.MarkFlagRequired("server")

	cloudDbGetCommand.Flags().StringP("host", "H", "", "The database host")
	cloudDbGetCommand.MarkFlagRequired("host")

	cloudDbGetCommand.Flags().StringP("db", "d", "", "The database name")
	cloudDbGetCommand.MarkFlagRequired("db")

}
