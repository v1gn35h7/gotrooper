package main

import (
	"flag"
	"fmt"

	"github.com/go-logr/zerologr"
	"github.com/goTrooper/pkg/constants"
	"github.com/goTrooper/pkg/kafka"
	"github.com/goTrooper/pkg/logging"
	"github.com/goTrooper/pkg/telemetry"
	"github.com/spf13/viper"
)

// CLI Flags
var (
	logger     zerologr.Logger
	verbrose   bool
	configPath string
)

func init() {

	flag.BoolVar(&verbrose, "verbrose", false, "Enabled verbrose mode.")
	flag.StringVar(&configPath, "config", "gotrroper.yml", "Config file path.")
	flag.Parse()

	// Set-up logger
	logger := logging.SetUpLogger()

	// Read config
	logger.Info("Reading config from file", "confi_path", configPath)
	viper.SetConfigName(constants.ConfigName) // name of config file (without extension)
	viper.SetConfigType(constants.ConfigType) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath)           // path to look for the config file in
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if verbrose {
		logger.Info("Setting verbrose mode on")
		logger.Info("Config read from file: \n", viper.AllSettings())
		fmt.Println(viper.AllSettings())
	}
}

func main() {
	fmt.Println("............................................  GoTrooper ................................................")

	logger.Info("GoTropper appliaction started.")

	//Bootstrap kafka client
	kclient := kafka.NewKafkaClient(logger)

	// Start telemetry collection
	tel := telemetry.NewTelemetryUtil(logger, kclient)
	tel.StartCollectingTelemetry()

}
