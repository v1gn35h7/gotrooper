package workers

import (
	"bufio"
	"context"
	"io"
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

func (hw *harvestWorker) StartHarvest(ctx context.Context) {
	// harvestInterval := viper.GetInt64("gotrooper.harvestInterval")
	// ticker := time.NewTicker(time.Second * time.Duration(harvestInterval))

	// hw.wg.Add(1)
	// // interval
	// go func() {
	// 	defer hw.wg.Done()

	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			hw.harvest()
	// 		case <-ctx.Done():
	// 			hw.logger.Info("Stopped Harvest workers...")
	// 			ticker.Stop()
	// 			hw.outputFile.File.Close()
	// 			return
	// 		}
	// 	}
	// }()
	hw.harvest()

}

func (hw *harvestWorker) harvest() {
	if hw.registryFile == nil {
		hw.registryFile = createRegFile(hw.logger)
	}

	//harvestPoint := readRegistryFile(hw.registryFile.File, hw.logger)

	finfo, _ := hw.outputFile.File.Stat()

	//fmt.Println(finfo.Size())

	if finfo.Size() != 0 {

		parsedData := make([]*pb.ShellScriptOutput, 0)

		br := bufio.NewReader(hw.outputFile.File)

		for {
			size, _, err := br.ReadLine()
			if err != nil {
				if viper.GetBool("verbose") == true {
					hw.logger.Error(err, "Failed to read file")
				}
				hw.resyncFile()
				br = bufio.NewReader(hw.outputFile.File)
				br.Discard(int(hw.fragmentSize))
				continue
			}

			s := strings.TrimSpace(string(size))
			hw.fragmentSize += int64(len([]byte(s)))

			//	_, err = br.Read(make([]byte, len([]byte(s))-1))

			if err != nil {
				if viper.GetBool("verbose") == true {
					hw.logger.Error(err, "Failed to read file")
				}
				hw.resyncFile()
				br = bufio.NewReader(hw.outputFile.File)
				br.Discard(int(hw.fragmentSize))
				continue
			}

			msgSize, _ := strconv.Atoi(s)
			hw.fragmentSize += int64(msgSize)

			pmsg := make([]byte, msgSize)

			r, err := io.ReadFull(br, pmsg)

			if err != nil {
				if viper.GetBool("verbose") == true {
					hw.logger.Error(err, "Failed to read file")
				}
				hw.resyncFile()
				br = bufio.NewReader(hw.outputFile.File)
				br.Discard(int(hw.fragmentSize))
				continue
			}

			if r == 0 {
				continue
			}

			//fmt.Println(string(pmsg))

			pbmsg := pb.ShellScriptOutput{}

			er := proto.Unmarshal(pmsg, &pbmsg)

			if er != nil {
				hw.logger.Error(er, "Failed to parse")
			}

			parsedData = append(parsedData, &pbmsg)

			if len(parsedData) > 10 {
				hw.shipDataFragment(parsedData)
				//clear(parsedData)
				parsedData = make([]*pb.ShellScriptOutput, 0)
			}
		}
	}
}

func (hw *harvestWorker) shipDataFragment(messages []*pb.ShellScriptOutput) {
	// Poll for scripts
	{
		c := pb.NewShellServiceClient(hw.grpcConc)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
		defer cancel()

		req := &pb.ShellFragmentRquest{}

		req.Outputs = messages

		r, err := c.SendFragment(ctx, req)
		if err != nil {
			hw.logger.Error(err, "could not send proto message")
		} else {
			hw.logger.Info("Response from gRPC server", "response", r.Awknowledgement)
		}
	}
}

func (hw *harvestWorker) resyncFile() {

	hdir, _ := os.UserHomeDir()

	hw.outputFile.File.Close()

	outFilePath := filepath.Join(hdir, "gotrooper.txtpb")

	ofile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		hw.logger.Error(err, "Failed to resync output file")
	}
	hw.outputFile.File = ofile
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

	//fmt.Println(fragment.String())

	fileContents := strings.Split(fragment.String(), "$$$#$$$")

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
