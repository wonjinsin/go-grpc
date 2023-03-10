package router

import (
	"context"
	v1 "phantom/controller/v1"
	"phantom/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

// InitHTTP ...
func InitHTTP(ctx context.Context, httpServer *runtime.ServeMux, port string, options []grpc.DialOption) {
	makeHTTPV1Route(ctx, httpServer, port, options)
}

func makeHTTPV1Route(ctx context.Context, httpServer *runtime.ServeMux, port string, options []grpc.DialOption) {
	v1.NewUserHTTPController(ctx, httpServer, port, options)
}
