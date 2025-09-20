package workers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pb"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
	"google.golang.org/protobuf/proto"
)

type harvestWorkerMock struct {
	logger       zerologr.Logger
	outputFile   *goshell.OutputFile
	fragmentData []byte
	fragmentSize int64
	prevFileSize int64
}

func NewHarvestWorkerMock(lgr zerologr.Logger, outputFile *goshell.OutputFile) *harvestWorkerMock {
	return &harvestWorkerMock{
		logger:       lgr,
		outputFile:   outputFile,
		fragmentData: make([]byte, 10240),
	}
}

func (hw *harvestWorkerMock) harvest() {
	// if hw.registryFile == nil {
	// 	hw.registryFile = createRegFile(hw.logger)
	// }

	//harvestPoint := readRegistryFile(hw.registryFile.File, hw.logger)

	finfo, _ := hw.outputFile.File.Stat()

	if finfo.Size() != 0 && finfo.Size() > hw.prevFileSize {
		hw.prevFileSize = finfo.Size()

		// for {
		// 	rb, err := hw.outputFile.File.Read(hw.fragmentData) //, int64(harvestPoint))

		// 	if err != nil {
		// 		if err.Error() == "EOF" {
		// 			break
		// 		}
		// 		hw.logger.Error(err, "Falied to read output file")
		// 	} else {
		// 		if rb == 0 {
		// 			break
		// 		}

		// 		hw.fragmentSize += int64(rb)

		// 		if hw.fragmentSize > 5 {
		// 			hw.parseFragmentFromBytes()
		// 		}

		// 	}
		// }

		// scanner := bufio.NewScanner(hw.outputFile.File)

		// scanner.Scan("") // Removed unused variable assignment

	}

}

// Parser bytes to proto
func (hw *harvestWorkerMock) parseFragmentFromBytes() {
	var fragment strings.Builder
	var parsedData []*pb.ShellScriptOutput

	for _, chunk := range hw.fragmentData {
		fragment.WriteByte(chunk)
	}

	fileContents := strings.Split(fragment.String(), "$$$#$$$")

	for _, msg := range fileContents {
		pbmsg := pb.ShellScriptOutput{}

		err := proto.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(msg), &pbmsg)

		if err != nil {
			hw.logger.Error(err, "error unmarchalling proto data")
			//break
		}
		parsedData = append(parsedData, &pbmsg)

	}

	fmt.Println(parsedData)
}

func TestHarvestWorker(t *testing.T) {

	// Setup App
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

	harvestWorker := NewHarvestWorkerMock(logger, outputFile)

	harvestWorker.harvest()

}
