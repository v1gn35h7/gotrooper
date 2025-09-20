package workers

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/go-logr/zerologr"
	"github.com/reugn/go-quartz/quartz"
	"github.com/v1gn35h7/gotrooper/internal/goshell"
	"github.com/v1gn35h7/gotrooper/pb"
	trooper "github.com/v1gn35h7/gotrooper/pkg/goTrooper"
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

func (e *executor) StartExecutors(ctx context.Context) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go execWorkers(ctx, rand.Int(), e)

}

func execWorkers(ctx context.Context, id int, exec *executor) {
	sem := make(chan struct{}, runtime.NumCPU())

	exec.wg.Add(1)
	defer exec.wg.Done()

	exec.logger.Info(fmt.Sprintf("Executor %d started...", id))

	// Reads from job queue and fires a goroutine to process the job
	for {
		select {
		case script := <-exec.jobQueue:
			exec.logger.Info("Job recevied", "script", script.Script)

			sem <- struct{}{}

			func() {
				exec.wg.Add(1)
				defer func() {
					exec.wg.Done()
					<-sem
				}()

				// Check local state
				_, ok := exec.jobs[script.Id]

				if !ok {

					// create jobs
					cronTrigger, _ := quartz.NewCronTrigger(script.Frequency)
					shellJob := trooper.NewShellJob(
						script.Script,
						exec.outputFile,
						exec.logger)
					jobDetail := quartz.NewJobDetail(shellJob, quartz.NewJobKey(script.Id))

					// register jobs to scheduler
					exec.scheduler.ScheduleJob(jobDetail, cronTrigger)

					// Add to state store
					exec.jobs[script.Id] = script

					exec.logger.Info("Job Submitted sucessfully...")
				}
			}()

		case <-ctx.Done():
			exec.logger.Info("Stopped executors...")
			close(sem)
			return
		}
	}

}
