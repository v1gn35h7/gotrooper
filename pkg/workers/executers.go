package workers

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pkg/logging"
	"github.com/v1gn35h7/gotrooper/pkg/shell"
)

type executor struct {
	logger     zerologr.Logger
	jobQueue   chan string
	wg         *sync.WaitGroup
	outputFile *goshell.OutputFile
}

func Executors(lgr zerologr.Logger, jq chan string, wgrp *sync.WaitGroup, outFile *goshell.OutputFile) *executor {

	return &executor{
		logger:     lgr,
		jobQueue:   jq,
		wg:         wgrp,
		outputFile: outFile,
	}
}

func (exec *executor) StartExecutors() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < runtime.NumCPU(); i++ {
		exec.wg.Add(1)

		go execWorker(ctx, i, exec)

	}

}

func execWorker(ctx context.Context, id int, exec *executor) {
	exec.logger.Info(fmt.Sprintf("Executor %d started...", id))
	for {
		select {
		case script := <-exec.jobQueue:
			// run script
			exec.logger.Info("Job recevied", "script", script)
			output, err := shell.ExecuteScript(script)

			if err != nil {
				exec.logger.Error(err, "Script execution failed", "scriptId", "")
			}

			logging.SaveOutputToLog(output, exec.outputFile)

		case <-ctx.Done():
			exec.wg.Done()
			return
		}
	}

}
