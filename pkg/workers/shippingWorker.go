package workers

import (
	"sync"
	"time"

	"github.com/go-logr/zerologr"
	"google.golang.org/grpc"
)

type shippingWorker struct {
	logger   zerologr.Logger
	grpcConc *grpc.ClientConn
	jobQueue chan string
	wg       *sync.WaitGroup
}

func ShippingWorker(lgr zerologr.Logger, conc *grpc.ClientConn, jq chan string, wgrp *sync.WaitGroup) *pollWorker {
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
func (pw *pollWorker) StartShipping(interval int64) {

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	quit := make(chan bool)
	pw.wg.Add(1)
	go func() {

		for {
			select {
			case <-ticker.C:
				//shipDataFragment(pw, quit)
			case <-quit:
				ticker.Stop()
				close(quit)
			}
		}
	}()

}
