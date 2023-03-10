package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"phantom/config"
	"phantom/controller"
	"phantom/repository"
	"phantom/service"
	"phantom/util"

	"github.com/dimiro1/banner"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var zlog *util.Logger

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[main] err[%s]", err.Error())
		os.Exit(1)
	}

	zlog.Infow("logger started")
	bannerInit()
}

func main() {
	phantom := config.Phantom

	repo, redis, err := repository.Init(phantom)
	if err != nil {
		zlog.Errorw("Error when Start repository: %v\n", err)
		os.Exit(1)
	}

	svc, err := service.Init(phantom, repo, redis)
	if err != nil {
		zlog.Errorw("Error when Start service: %v\n", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	controller.Init(server, svc)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", phantom.GetString("port")))
	if err != nil {
		zlog.Errorw("Error when Start listen port: %v\n", err)
		os.Exit(1)
	}

	httpServer := runtime.NewServeMux()
	controller.InitHTTPController(httpServer, fmt.Sprintf("localhost:%s", phantom.GetString("port")))
	log.Printf("start HTTP server on %s port", fmt.Sprintf("localhost:%s", phantom.GetString("httpPort")))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", phantom.GetString("httpPort")), httpServer); err != nil {
		zlog.Errorw("failed to serve http server: %v", err)
		os.Exit(1)
	}

	log.Printf("start gRPC server on %s port", phantom.GetString("port"))
	if err := server.Serve(lis); err != nil {
		zlog.Errorw("failed to serve: %v", err)
		os.Exit(1)
	}
}

func bannerInit() {
	isEnabled := true
	isColorEnabled := true
	in, err := os.Open("banner.txt")
	if in == nil || err != nil {
		os.Exit(1)
	}

	banner.Init(os.Stdout, isEnabled, isColorEnabled, in)
}
