package cli

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/internal/client"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
	"github.com/v1gn35h7/gotrooper/pkg/workers"
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

			// Setup output file
			outputFilePath := viper.GetString("outputFile")

			if outputFilePath == "" {
				homeDir, _ := os.UserHomeDir()
				outputFilePath = filepath.Join(homeDir, "gotrooper.log")

			}

			file, err := os.OpenFile(outputFilePath, os.O_CREATE, os.ModeAppend)

			if err != nil {
				logger.Error(err, "Failed to create output file")
			}

			outputFile := &goshell.OutputFile{
				File: file,
			}

			// Set-up gRPC client
			conc := client.SetupGrpcClient(logger)
			defer conc.Close()

			// Create job quee
			jobQueue := make(chan string, 100)

			var wg sync.WaitGroup
			defer wg.Done()

			//Start executor workers
			workers.Executors(logger, jobQueue, &wg, outputFile).StartExecutors()

			// Start polling go routine
			pollInterval := viper.GetInt64("goshell.refreshInterval")
			workers.PollWorker(logger, conc, jobQueue, &wg).StartPolling(pollInterval)

			// Wait for all go routines to complete
			wg.Wait()

		},
	}
	return startCmd
}
