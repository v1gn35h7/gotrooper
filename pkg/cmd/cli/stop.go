package cli

import (
	"github.com/spf13/cobra"
)

func NewStopCommand() *cobra.Command {
	var stopCmd = &cobra.Command{
		Use:   "version",
		Short: "goshellctl stope",
		Long:  "Stopa Goshell version",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return stopCmd
}
