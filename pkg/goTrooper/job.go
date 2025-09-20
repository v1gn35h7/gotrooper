package trooper

import (
	"context"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	gpb "github.com/v1gn35h7/gotrooper/internal/pb"
	"github.com/v1gn35h7/gotrooper/pkg/shell"
)

type ShellJob struct {
	Command    string
	logger     zerologr.Logger
	outputFile *goshell.OutputFile
}

func (sh *ShellJob) Execute(ctx context.Context) error {
	output, err := shell.ExecuteScript(sh.Command)

	if err != nil {
		sh.logger.Error(err, "Script execution failed", "scriptId", "")
		return err
	}

	//logging.SaveOutputToLog(output, exec.outputFile)
	err = gpb.SaveOutputToPb(output, sh.outputFile)

	if err != nil {
		sh.logger.Error(err, "Error saving the output")
		return err
	}
	return nil
}

func (sh ShellJob) Description() string {
	return "Shell Job"
}

func NewShellJob(cmd string, of *goshell.OutputFile, logger zerologr.Logger) *ShellJob {
	return &ShellJob{
		Command:    cmd,
		outputFile: of,
		logger:     logger,
	}
}
