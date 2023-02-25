package controller

import (
	"fmt"
	"log"
	"net"
	"os"
	"phantom/util"

	"google.golang.org/grpc"
)

var zlog *util.Logger

type phantomStatus struct {
	TRID       string      `json:"trid"`
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Init ...
func Init(port string) (server *grpc.Server, err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		zlog.Errorw("failed to listen: %v", err)
		return nil, err
	}

	grpcServer := grpc.NewServer()
	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		zlog.Errorw("failed to serve: %s", err)
		return nil, err
	}

	return server, nil
}
