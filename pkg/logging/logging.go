package logging

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
)

var (
	lfile *os.File
	log   logr.Logger
)

func init() {
	homeDir, _ := os.UserHomeDir()

	// Set-up logger
	f, err := os.OpenFile(filepath.Join(homeDir, "gotrooper.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err, "Failed to create log file")
		os.Exit(1)
	}

	lfile = f
}

func Close() {
	lfile.Close()
}

// Setups logger instance
func Logger() logr.Logger {

	if log.GetSink() != nil {
		return log
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zerologr.SetMaxV(1)

	zl := zerolog.New(lfile)
	zl = zl.With().Caller().Timestamp().Logger()
	log = zerologr.New(&zl)
	return log
}

func SaveOutputToLog(output string, outputFile *goshell.OutputFile) error {
	outputFile.Mutex.Lock()
	defer func() {
		outputFile.Mutex.Unlock()
	}()

	hn, _ := os.Hostname()

	scriptOutput := goshell.ScriptOutput{
		AgentId:  uuid.NewString(),
		HostName: hn,
		ScriptId: uuid.NewString(),
		Output:   output,
	}

	out, _ := json.Marshal(scriptOutput)
	_, err := outputFile.File.Write(out)
	return err
}
