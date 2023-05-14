package workers

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/pkg/shell"
)

type executor struct {
	logger   zerologr.Logger
	jobQueue chan string
	wg       *sync.WaitGroup
}

func Executors(lgr zerologr.Logger, jq chan string, wgrp *sync.WaitGroup) *executor {

	return &executor{
		logger:   lgr,
		jobQueue: jq,
		wg:       wgrp,
	}
}

func (exec *executor) StartExecutors() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < runtime.NumCPU(); i++ {
		exec.wg.Add(1)

		go execWorker(ctx, i, exec.wg, exec.logger, exec.jobQueue)

	}

}

func execWorker(ctx context.Context, id int, wg *sync.WaitGroup, logger zerologr.Logger, jobQueue chan string) {
	logger.Info(fmt.Sprintf("Executor %d started...", id))
	for {
		select {
		case script := <-jobQueue:
			// run script
			logger.Info("Job recevied", "script", script)
			shell.ExecuteScript(script)

		case <-ctx.Done():
			wg.Done()
			return
		}
	}

}
