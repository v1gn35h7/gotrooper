package cli

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

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

			viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

			homeDir, _ := os.UserHomeDir()

			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")
			defer logging.Close()

			// Setup output file
			outputFilePath := viper.GetString("outputFile")

			if outputFilePath == "" {
				outputFilePath = filepath.Join(homeDir, "gotrooper.txtpb")
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

			// Setup clean up
			ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

			// Create job quee
			jobQueue := make(chan *pb.ShellScript, 100)

			var wg sync.WaitGroup

			// Setup internal store
			//store := store.NewTrooperStore()

			// Starts scheduuler
			trooper_schedlr := scheduler.TooperScheduler(ctx)

			//Starts executor workers
			workers.Executors(logger, jobQueue, &wg, outputFile, trooper_schedlr).StartExecutors(ctx)

			// Starts polling worker
			pollInterval := viper.GetInt64("gotrooper.goshell.refreshInterval")
			workers.PollWorker(logger, conc, jobQueue, &wg).StartPolling(ctx, pollInterval)

			// Starts upload worker
			regFilePath := filepath.Join(homeDir, "gotrooper_registry.log")
			regfile, err := os.OpenFile(regFilePath, os.O_CREATE, 0777)
			if err != nil {
				logger.Error(err, "Failed to open registry file")
			}
			defer regfile.Close()

			registryFile := &goshell.RegistryFile{
				File: regfile,
			}

			workers.NewHarvestWorker(logger, conc, &wg, outputFile, registryFile).StartHarvest(ctx)

			// Wait for all go routines to complete
			wg.Wait()

		},
	}
	return startCmd
}
