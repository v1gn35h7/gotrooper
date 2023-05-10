package cli

import (
	"github.com/spf13/cobra"
	"github.com/v1gn35h7/gotrooper/pkg/kafka"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
	"github.com/v1gn35h7/gotrooper/pkg/telemetry"
)

func NewStartCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "gotrooper start",
		Long:  "Starts gotrooper cli process",
		Run: func(cmd *cobra.Command, args []string) {
			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			//Bootstrap kafka client
			kclient := kafka.NewKafkaClient(logger)

			// Start telemetry collection
			tel := telemetry.NewTelemetryUtil(logger, kclient)
			tel.StartCollectingTelemetry()

		},
	}
	return startCmd
}
