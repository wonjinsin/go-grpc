package v1

import (
	"context"
	"log"
	"phantom/model"
	"phantom/service"
	"phantom/util"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
)

// UserServer ...
type UserServer struct {
	UserControllerServer
	userSvc service.UserService
}

// NewUserController ...
func NewUserController(server *grpc.Server, userSvc service.UserService) {
	RegisterUserControllerServer(server, &UserServer{
		userSvc: userSvc,
	})
}

// NewUserHTTPController ...
func NewUserHTTPController(ctx context.Context, httpServer *runtime.ServeMux, port string, options []grpc.DialOption) {
	if err := RegisterUserControllerHandlerFromEndpoint(
		ctx,
		httpServer,
		port,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}
}

// GetUser ...
func (u UserServer) GetUser(ctx context.Context, req *GetUserRequest) (user *model.User, err error) {
	zlog.With(ctx).Infow("[New request]", "req", req)
	intCtx, cancel := context.WithTimeout(ctx, util.CtxTimeOut)
	defer cancel()

	uid := req.Uid
	if _, err = uuid.Parse(uid); err != nil {
		zlog.With(intCtx).Warnw("ID is not valid", "uid", uid, "err", err)
		return nil, err
	}

	if user, err = u.userSvc.GetUser(intCtx, uid); err != nil {
		zlog.With(intCtx).Warnw("UserSvc GetUser failed", "uid", uid, "err", err)
		return nil, err
	}

	return user, nil
}
