package client

import (
	"os"
	"time"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func SetupGrpcClient(logger logr.Logger) *grpc.ClientConn {
	grpcAddr := viper.GetString("goshell.host")

	conc, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))

	if err != nil {
		logger.Error(err, "grpc connection failed")
		os.Exit(1)
	}

	return conc
}
