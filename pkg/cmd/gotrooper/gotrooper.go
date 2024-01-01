package gotrooper

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v1gn35h7/gotrooper/internal/config"
	"github.com/v1gn35h7/gotrooper/pkg/cmd/cli"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
)

var (
	configPath string
	verbose    bool
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Gotrooper",
		Short: "Gotrooper service",
		Long:  "Gotrooper service starts Gotrooper command line services",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			printLogo()

			// Init read config
			fmt.Println("Config path set to: ", configPath)
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
	fmt.Println("############################################################################################################################################")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("        \"\"\"\"\"\"\"                                   ")
	fmt.Println("       \"\"       \"\"                                  ")
	fmt.Println("      \"\"                                              ")
	fmt.Println("      \"\"                                               ")
	fmt.Println("      \"\"                \"\"\"\"\"\"  \"\"\"\"\"\"  \"\"\"\"\"\"  \"\"\"\"\"  \"\"\"\"\"  \"\"\"\"\"  \"\"\"\"\"   \"\"\"\"\"\" ")
	fmt.Println("      \"\"                \"        \"       \"       \"    	 \"							  \"	 \"	 \"			   \"		\" ")
	fmt.Println("      \"\"         \"\"   \"        \"       \"       \"        \"  \" 	 \"	 \"		 \"	  \"	 \"	 \"			   \"		\" ")
	fmt.Println("       \"\"        \"\"   \"        \"       \"       \" \"\"\"\"   						  \"\"\"\"	 \"\"\"\"\"    \" \"\"\"\" ")
	fmt.Println("         \"\"      \"\"   \"        \"       \"       \"     	\"	 \"		 \"	 \"		 \"	  \"		 \"			   \"		\" ")
	fmt.Println("           \"\"\"\"\"\"   \"\"\"\"\"\"       \"       \"    	\"	 \"\"\"\"\"  \"\"\"\"\"   \"		 \"\"\"\"\"	   \"       \" ")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("############################################################################################################################################")
}
