/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "shcli",
		Short: "CLI tool for interacting with the Sitehost api",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".shcli.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			// Determine the naming convention of the flags when represented in the config file
			configName := f.Name

			// Apply the viper config value to the flag when the flag is not set and viper has a value
			if !f.Changed && viper.IsSet(configName) {
				rootCmd.PersistentFlags().Set(f.Name, viper.GetString(configName))
			}
		})
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shcli.json)")
	rootCmd.PersistentFlags().StringP("apiKey", "", "", "Sitehost api key")
	rootCmd.PersistentFlags().StringP("clientId", "", "", "Sitehost client id")
	rootCmd.MarkPersistentFlagRequired("apiKey")
	rootCmd.MarkPersistentFlagRequired("clientId")

	viper.BindPFlag("apiKey", rootCmd.PersistentFlags().Lookup("apiKey"))
	viper.BindPFlag("clientId", rootCmd.PersistentFlags().Lookup("clientId"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}