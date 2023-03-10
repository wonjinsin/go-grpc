package controller

import (
	"context"
	"log"
	"os"
	"phantom/router"
	"phantom/service"
	"phantom/util"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
func Init(server *grpc.Server, svc *service.Service) {
	router.Init(server, svc)
}

// InitHTTPController ...
func InitHTTPController(httpServer *runtime.ServeMux, port string) {
	ctx := context.Background()
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	router.InitHTTP(ctx, httpServer, port, options)
}
