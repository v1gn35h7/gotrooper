package logging

import (
	"encoding/json"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
)

// Setups logger instance
func Logger() logr.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zerologr.SetMaxV(1)

	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()
	var log logr.Logger = zerologr.New(&zl)
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
