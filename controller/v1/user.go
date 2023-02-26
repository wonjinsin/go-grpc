package v1

import (
	"context"
	"phantom/model"
	"phantom/service"
	"phantom/util"

	"github.com/google/uuid"
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
