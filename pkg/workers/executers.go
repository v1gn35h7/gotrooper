package workers

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/go-logr/zerologr"
	"github.com/reugn/go-quartz/quartz"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	gpb "github.com/v1gn35h7/gotrooper/internal/pb"
	"github.com/v1gn35h7/gotrooper/pb"
	"github.com/v1gn35h7/gotrooper/pkg/shell"
)

type executor struct {
	logger     zerologr.Logger
	jobQueue   chan *pb.ShellScript
	wg         *sync.WaitGroup
	outputFile *goshell.OutputFile
	scheduler  quartz.Scheduler
	jobs       map[string]*pb.ShellScript
}

func Executors(lgr zerologr.Logger, jq chan *pb.ShellScript, wgrp *sync.WaitGroup, outFile *goshell.OutputFile, scheduler quartz.Scheduler) *executor {

	return &executor{
		logger:     lgr,
		jobQueue:   jq,
		wg:         wgrp,
		outputFile: outFile,
		scheduler:  scheduler,
		jobs:       make(map[string]*pb.ShellScript),
	}
}

func (exec *executor) StartExecutors() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < 1; i++ {
		exec.wg.Add(1)

		go execWorker(ctx, i, exec)

	}

}

func execWorker(ctx context.Context, id int, exec *executor) {
	exec.logger.Info(fmt.Sprintf("Executor %d started...", id))
	for {
		select {
		case script := <-exec.jobQueue:
			exec.logger.Info("Job recevied", "script", script.Script)

			// Check local state
			_, ok := exec.jobs[script.Id]

			if !ok {

				// create jobs
				cronTrigger, _ := quartz.NewCronTrigger(script.Frequency)
				functionJob := quartz.NewFunctionJob(func(_ context.Context) (string, error) {
					output, err := shell.ExecuteScript(script.Script)

					if err != nil {
						exec.logger.Error(err, "Script execution failed", "scriptId", "")
						return "", err
					}

					//logging.SaveOutputToLog(output, exec.outputFile)
					err = gpb.SaveOutputToPb(output, exec.outputFile)

					if err != nil {
						exec.logger.Error(err, "Error saving the output")
						return "", err
					}
					return output, nil
				})

				// register jobs to scheduler
				exec.scheduler.ScheduleJob(ctx, functionJob, cronTrigger)

				// Add to state store
				exec.jobs[script.Id] = script
			}

		case <-ctx.Done():
			exec.wg.Done()
			return
		}
	}

}
