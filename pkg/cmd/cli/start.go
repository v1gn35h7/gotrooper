package cli

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/internal/client"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pb"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
	"github.com/v1gn35h7/gotrooper/pkg/scheduler"
	"github.com/v1gn35h7/gotrooper/pkg/workers"
)

func NewStartCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "gotrooper start",
		Long:  "Starts gotrooper cli application",
		Run: func(cmd *cobra.Command, args []string) {
			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			homeDir, _ := os.UserHomeDir()

			// Setup output file
			outputFilePath := viper.GetString("outputFile")

			if outputFilePath == "" {
				outputFilePath = filepath.Join(homeDir, "gotrooper.log")
			}

			file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			defer file.Close()

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
			jobQueue := make(chan *pb.ShellScript, 100)

			var wg sync.WaitGroup
			defer wg.Done()

			// Setup internal store
			//store := store.NewTrooperStore()

			// Starts scheduuler
			trooper_schedlr, scheduler_cancel := scheduler.TooperScheduler()
			defer scheduler_cancel()

			//Starts executor workers
			workers.Executors(logger, jobQueue, &wg, outputFile, trooper_schedlr).StartExecutors()

			// Starts polling worker
			pollInterval := viper.GetInt64("gotrooper.goshell.refreshInterval")
			workers.PollWorker(logger, conc, jobQueue, &wg).StartPolling(pollInterval)

			// Starts upload worker]
			regFilePath := filepath.Join(homeDir, "gotrooper_registry.log")
			regfile, err := os.OpenFile(regFilePath, os.O_CREATE, 0777)
			if err != nil {
				logger.Error(err, "Failed to open registry file")
			}
			defer regfile.Close()

			registryFile := &goshell.RegistryFile{
				File: regfile,
			}

			workers.NewHarvestWorker(logger, conc, &wg, outputFile, registryFile).StartHarvest()

			// Wait for all go routines to complete
			wg.Wait()

		},
	}
	return startCmd
}
