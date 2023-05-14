package workers

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/pb"
	"github.com/v1gn35h7/gotrooper/pkg/shell"
	"google.golang.org/grpc"
)

type pollWorker struct {
	logger   zerologr.Logger
	grpcConc *grpc.ClientConn
}

func PollWorker(lgr zerologr.Logger, conc *grpc.ClientConn) *pollWorker {
	return &pollWorker{
		logger:   lgr,
		grpcConc: conc,
	}
}

/*
* https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals
 */
func (pw *pollWorker) StartPolling(interval int64) {

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	quit := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

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

	// Wait for all goroutine to complete
	wg.Wait()

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
				shell.ExecuteScript(v.Script)
			}
		}
	}
}
