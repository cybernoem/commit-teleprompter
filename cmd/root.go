package cmd

import (
	"fmt"
	"os"

	"github.com/cybernoem/commit-teleprompter/cmd/issue"
	"github.com/cybernoem/commit-teleprompter/cmd/message"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "c-t",
	Long:    "Commit Teleprompter (c-t) is a CLI tool that was created to help writing commit messages",
	Example: "c-t list",

	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(issue.Cmd)
	RootCmd.AddCommand(message.Cmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetDefault("IssueKey", "")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.ct-prompt")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Not found")
			viper.SafeWriteConfig()
		} else { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}
