package workers

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const regFilePath = "registry"

type harvestWorker struct {
	logger       zerologr.Logger
	grpcConc     *grpc.ClientConn
	wg           *sync.WaitGroup
	registryFile *goshell.RegistryFile
	outputFile   *goshell.OutputFile
	fragmentData []byte
	fragmentSize int64
	prevFileSize int64
}

func NewHarvestWorker(lgr zerologr.Logger, conc *grpc.ClientConn, wgrp *sync.WaitGroup, outputFile *goshell.OutputFile, registryFile *goshell.RegistryFile) *harvestWorker {
	return &harvestWorker{
		logger:       lgr,
		grpcConc:     conc,
		wg:           wgrp,
		outputFile:   outputFile,
		registryFile: registryFile,
		fragmentData: make([]byte, 10240),
	}
}

func (hw *harvestWorker) StartHarvest() {
	harvestInterval := viper.GetInt64("gotrooper.harvestInterval")
	ticker := time.NewTicker(time.Second * time.Duration(harvestInterval))
	quit := make(chan bool, 1)
	hw.wg.Add(1)

	// interval
	go func() {

		for {
			select {
			case <-ticker.C:
				hw.harvest()
			case <-quit:
				ticker.Stop()
				hw.wg.Done()
				close(quit)
			}
		}
	}()

}

func (hw *harvestWorker) harvest() {
	if hw.registryFile == nil {
		hw.registryFile = createRegFile(hw.logger)
	}

	//harvestPoint := readRegistryFile(hw.registryFile.File, hw.logger)

	finfo, _ := hw.outputFile.File.Stat()

	if finfo.Size() != 0 && finfo.Size() > hw.prevFileSize {
		hw.prevFileSize = finfo.Size()

		for {
			rb, err := hw.outputFile.File.Read(hw.fragmentData) //, int64(harvestPoint))

			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				hw.logger.Error(err, "Falied to read output file")
			} else {
				if rb == 0 {
					break
				}

				hw.fragmentSize += int64(rb)

				if hw.fragmentSize > 100 {
					hw.shipDataFragment()
				}

			}
		}
	}

	// Sent to server
	//if rb ==

}

func (hw *harvestWorker) shipDataFragment() {
	// Poll for scripts
	{
		c := pb.NewShellServiceClient(hw.grpcConc)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
		defer cancel()

		req := &pb.ShellFragmentRquest{}
		//req.Outputs = append(req.Outputs)
		outputPorots := hw.parseFragmentFromBytes()

		req.Outputs = outputPorots

		r, err := c.SendFragment(ctx, req)
		if err != nil {
			hw.logger.Error(err, "could not send proto message")
		} else {
			hw.logger.Info("Response from gRPC server", "response", r.Awknowledgement)
		}
	}
}

func createRegFile(logger zerologr.Logger) *goshell.RegistryFile {
	homeDir, _ := os.UserHomeDir()
	regFilePath := filepath.Join(homeDir, "registry")
	regfile, err := os.OpenFile(regFilePath, os.O_RDWR, os.ModeExclusive)

	if err != nil {
		logger.Error(err, "Failed to registry output file")
	}

	return &goshell.RegistryFile{
		File: regfile,
	}
}

// Read registry file
func readRegistryFile(regfile *os.File, logger zerologr.Logger) int64 {
	var harvestPoint int64
	var byt []byte
	_, err := regfile.Read(byt)

	if err != nil {
		logger.Error(err, "Failed to read registry file")
		harvestPoint = 0
	} else {
		x, _ := strconv.Atoi(string(byt))
		harvestPoint = int64(x)
	}

	return harvestPoint
}

// Parser bytes to proto
func (hw *harvestWorker) parseFragmentFromBytes() []*pb.ShellScriptOutput {
	var fragment strings.Builder
	var parsedData []*pb.ShellScriptOutput

	for _, chunk := range hw.fragmentData {
		fragment.WriteByte(chunk)
	}

	fileContents := strings.Split(fragment.String(), "\n\n\n")

	for _, msg := range fileContents {
		pbmsg := pb.ShellScriptOutput{}

		err := proto.Unmarshal([]byte(msg), &pbmsg)

		if err != nil {
			hw.logger.Error(err, "error unmarchalling proto data")
			//break
		}
		parsedData = append(parsedData, &pbmsg)

	}

	return parsedData
}
