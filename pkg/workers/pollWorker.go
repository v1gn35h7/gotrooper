package workers

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/pb"
	"google.golang.org/grpc"
)

type pollWorker struct {
	logger   zerologr.Logger
	grpcConc *grpc.ClientConn
	jobQueue chan string
	wg       *sync.WaitGroup
}

func PollWorker(lgr zerologr.Logger, conc *grpc.ClientConn, jq chan string, wgrp *sync.WaitGroup) *pollWorker {
	return &pollWorker{
		logger:   lgr,
		grpcConc: conc,
		jobQueue: jq,
		wg:       wgrp,
	}
}

/*
* https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals
 */
func (pw *pollWorker) StartPolling(interval int64) {

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	quit := make(chan bool)
	pw.wg.Add(1)
	go func() {

		for {
			select {
			case <-ticker.C:
				getScripts(pw, quit)
			case <-quit:
				ticker.Stop()
				close(quit)
			}
		}
	}()

}

func getScripts(pw *pollWorker, quit chan bool) {
	// Poll for scripts
	{
		c := pb.NewShellServiceClient(pw.grpcConc)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
		defer cancel()

		hostName, _ := os.Hostname()
		r, err := c.GetScripts(ctx, &pb.ShellRequest{AgentId: hostName})
		if err != nil {
			pw.logger.Error(err, "could not send proto message")
			quit <- true
		} else {
			pw.logger.Info("Response from gRPC server", "response", r.Scripts)
			for _, v := range r.Scripts {
				pw.jobQueue <- v.Script
			}
		}
	}
}
