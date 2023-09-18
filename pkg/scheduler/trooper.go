package scheduler

import (
	"context"

	"github.com/reugn/go-quartz/quartz"
)

func TooperScheduler() (quartz.Scheduler, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	// create scheduler
	sched := quartz.NewStdScheduler()

	// async start scheduler
	sched.Start(ctx)

	// wait for all workers to exit
	// sched.Wait(ctx)

	return sched, cancel
}
