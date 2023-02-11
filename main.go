package main

import (
	"fmt"
	"log"
	"os"
	"phantom/config"
	"phantom/controller"
	"phantom/repository"
	"phantom/service"
	"phantom/util"

	"github.com/dimiro1/banner"
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

	server, err := controller.Init(phantom.GetString("port"))
	if err != nil {
		fmt.Printf("Error when Start server: %v\n", err)
		os.Exit(1)
	}

	repo, redis, err := repository.Init(phantom)
	if err != nil {
		fmt.Printf("Error when Start repository: %v\n", err)
		os.Exit(1)
	}

	svc, err := service.Init(phantom, repo, redis)
	if err != nil {
		fmt.Printf("Error when Start service: %v\n", err)
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
