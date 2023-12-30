package issue

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SetCmd = &cobra.Command{
	Use:     "set [issue key]",
	Short:   "Sets the current issue key",
	Long:    "Sets the current issue key. The key is kept generic to support any issue tracker system.",
	Example: "c-t issue set GIT-432",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an issue key")
			return
		}
		issueKey := args[0]

		if issueKey == "" {
			fmt.Println("Please provide an issue key")
			return
		}

		viper.Set("IssueKey", issueKey)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Error saving config file:", err)
		} else {
			fmt.Println("Config file saved successfully.")
		}
		fmt.Printf("Issue key set to %s\n", viper.GetString("IssueKey"))
		return
	},
}
