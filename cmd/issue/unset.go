package issue

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UnsetCmd = &cobra.Command{
	Use:     "unset [issue key]",
	Short:   "Unsets the current issue key",
	Example: "c-t issue unset",
	Aliases: []string{"u"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("IssueKey", "")
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Error saving config file:", err)
		} else {
			fmt.Println("Config file saved successfully.")
		}
	},
}
