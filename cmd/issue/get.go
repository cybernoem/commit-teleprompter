package issue

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get current issue key",
	Long:    "Get which issue is selected to be used in commit templates",
	Example: "c-t issue get",
	Aliases: []string{"g", "get"},
	Run: func(cmd *cobra.Command, args []string) {
		var currentKey = viper.Get("IssueKey")
		if currentKey != "" {
			fmt.Printf("The current issue key is set to %s\n", currentKey)
		} else {
			fmt.Printf("The current issue key is not set.\nTo set it: c-t issue set [ISSUE_KEY]\n")
		}
	},
}
