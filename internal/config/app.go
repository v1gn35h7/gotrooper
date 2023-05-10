package config

import (
	"fmt"

	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/pkg/constants"
)

func ReadConfig(configPath string, logger zerologr.Logger) {
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

	fmt.Println(viper.AllSettings())
}
