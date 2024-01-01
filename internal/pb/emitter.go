package pb

import (
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	pbf "github.com/v1gn35h7/gotrooper/pb"
	"google.golang.org/protobuf/proto"
)

func SaveOutputToPb(output string, outputFile *goshell.OutputFile) error {
	outputFile.Mutex.Lock()
	defer func() {
		outputFile.Mutex.Unlock()
	}()

	hn, _ := os.Hostname()

	shellOut := &pbf.ShellScriptOutput{
		AgentId:  viper.GetString("gotrooper.hostId"),
		HostName: hn,
		Output:   output,
		ScriptId: uuid.NewString(),
	}

	out, _ := proto.Marshal(shellOut)

	size := len(out)

	// Add size
	_, err := outputFile.File.Write([]byte(strconv.Itoa(size)))
	if err != nil {
		return err
	}

	_, err = outputFile.File.Write([]byte("\n"))
	if err != nil {
		return err
	}

	_, err = outputFile.File.Write(out)
	if err != nil {
		return err
	}

	return err
}
