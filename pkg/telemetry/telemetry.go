package telemetry

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-logr/zerologr"
	"github.com/v1gn35h7/gotrooper/pb"
	"golang.org/x/sys/windows"
	"google.golang.org/grpc"
)

// unsafe.Sizeof(windows.ProcessEntry32{})
const processEntrySize = 568

type TelemetryUtil struct {
	logger          zerologr.Logger
	processSnapShot map[string]string
	grpcConc        *grpc.ClientConn
}

func NewTelemetryUtil(logr zerologr.Logger, conc *grpc.ClientConn) *TelemetryUtil {
	return &TelemetryUtil{
		logger:          logr,
		processSnapShot: make(map[string]string),
		grpcConc:        conc,
	}
}

/*
* Single threaded telemetry collection
 */
func (tele *TelemetryUtil) StartCollectingTelemetry() {
	for {
		tele.enumProcessess()
	}
}

func (tele *TelemetryUtil) collectWindowsProcess() {
	//for {
	//enumProcessess()
	//}
}

/*
* Took from
* https://stackoverflow.com/questions/11356264/list-of-currently-running-process-in-golang-windows-version
 */
func (tele *TelemetryUtil) enumProcessess() {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	defer windows.CloseHandle(h)
	if e != nil {
		panic(e)
	}
	p := windows.ProcessEntry32{Size: processEntrySize}
	for {
		e := windows.Process32Next(h, &p)
		pid := strconv.FormatUint(uint64(p.ProcessID), 10)
		if e != nil {
			break
		}
		s := windows.UTF16ToString(p.ExeFile[:])
		if tele.processSnapShot[pid] == "" {
			tele.processSnapShot[pid] = s
			tele.submitTrooperEvent(s)
		}
	}
}

func (tele *TelemetryUtil) submitTrooperEvent(event string) {

	// TODO: Send proto bufs to server
	c := pb.NewShellServiceClient(tele.grpcConc)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	hostName, _ := os.Hostname()
	r, err := c.GetScripts(ctx, &pb.ShellRequest{AgentId: hostName})
	if err != nil {
		tele.logger.Error(err, "could not send proto message")
	} else {
		tele.logger.Info("Response from gRPC server", "response", r.Scripts)
	}

}
