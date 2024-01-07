package issue

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:           "issue",
	Short:         "Manage the current selected issue",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	Cmd.AddCommand(GetCmd)
	Cmd.AddCommand(SetCmd)
	Cmd.AddCommand(UnsetCmd)
}
