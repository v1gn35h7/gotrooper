package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewVersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Gotrooper version",
		Long:  "Prints gotrooper version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Gotrooper ", viper.GetString("gotrooper.version"))
		},
	}
	return versionCmd
}
