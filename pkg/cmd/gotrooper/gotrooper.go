package gotrooper

import (
	"fmt"
	"log/slog"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
	"github.com/v1gn35h7/gotrooper/internal/config"
	"github.com/v1gn35h7/gotrooper/pkg/cmd/cli"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
)

var (
	configPath string
	verbose    bool
	ascii      = figlet4go.NewAsciiRender()
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Gotrooper",
		Short: "Gotrooper service",
		Long:  "Gotrooper service starts Gotrooper command line services",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			printLogo()

			// Init read config
			slog.Info("Config path set", slog.String("configPath", configPath))
			config.ReadConfig(configPath, logging.Logger())
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "config file path")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", true, "verbose mode")

	// Add sub commands
	rootCmd.AddCommand(cli.NewVersionCommand())
	rootCmd.AddCommand(cli.NewStartCommand())
	rootCmd.AddCommand(cli.NewStopCommand())
	rootCmd.AddCommand(cli.NewDebugCommand())

	return rootCmd
}

func printLogo() {
	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorMagenta,
	}
	options.FontName = "larry3d"

	renderStr, _ := ascii.RenderOpts("GoTrooper", options)

	fmt.Print(renderStr)
}
