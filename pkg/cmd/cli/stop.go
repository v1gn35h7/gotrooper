package cli

import (
	"github.com/spf13/cobra"
)

func NewStopCommand() *cobra.Command {
	var stopCmd = &cobra.Command{
		Use:   "version",
		Short: "gotrooper stope",
		Long:  "Stopa gotrooper version",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return stopCmd
}
