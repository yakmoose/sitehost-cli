/*
Copyright © 2023 John Lennard <john@yakmoo.se>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"errors"

	"github.com/sitehostnz/gosh/pkg/api"
	"github.com/sitehostnz/gosh/pkg/api/cloud/image"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloudImageListCmd represents the list command
var cloudImageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List images",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := image.New(api.NewClient(viper.GetString("apiKey"), viper.GetString("clientId")))

		imageResponse, err := client.List(context.Background())
		if err != nil {
			return err
		}

		format := cmd.Flag("format").Value.String()
		if format == "json" {
			json, err := json.MarshalIndent(imageResponse.Return.Images, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(json))
		} else if format == "text" {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 4, 4, ' ', 0)
			fmt.Fprintln(w, "Image Id\tImage Label\tImage Code\tVersion Count\tContainer Count")
			for _, image := range imageResponse.Return.Images {
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\n", image.ID, image.Label, image.Code, image.VersionCount, image.ContainerCount)
			}

			fmt.Fprintln(w)
		} else {
			/// how to error out.
			return errors.New("unsupported format, please choose text or json")
		}

		return nil
	},
}

func init() {
	cloudImageCommand.AddCommand(cloudImageListCmd)
}
