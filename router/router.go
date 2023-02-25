package router

import (
	v1 "phantom/controller/v1"
	"phantom/service"

	"google.golang.org/grpc"
)

// Init ...
func Init(server *grpc.Server, svc *service.Service) {
	makeV1Route(server, svc)
}

func makeV1Route(server *grpc.Server, svc *service.Service) {
	makeV1UserRoute(server, svc)
}

func makeV1UserRoute(server *grpc.Server, svc *service.Service) {
	v1.NewUserController(server, svc.User)
}
