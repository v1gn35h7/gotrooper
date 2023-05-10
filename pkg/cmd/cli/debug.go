package cli

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

func NewDebugCommand() *cobra.Command {
	var debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "goshellctl debug",
		Long:  "Debug Util for GoshellCtl",
		Run: func(cmd *cobra.Command, args []string) {
			debug.PrintStack()
		},
	}
	return debugCmd
}
