package v1

import (
	"context"
	"net/http"
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
func NewUserController(grpcServer *grpc.Server, userSvc service.UserService) {
	RegisterUserControllerServer(grpcServer, &UserServer{})
}

// GetUser ...
func (u UserServer) GetUser(ctx context.Context, req GetUserRequest) (err error) {
	ctx := c.Request().Context()
	uid := c.Param("uid")
	zlog.With(ctx).Infow("[New request]", "uid", uid)
	intCtx, cancel := context.WithTimeout(ctx, util.CtxTimeOut)
	defer cancel()

	if _, err = uuid.Parse(uid); err != nil {
		zlog.With(intCtx).Warnw("ID is not valid", "uid", uid, "err", err)
		return response(c, http.StatusBadRequest, "User is not valid")
	}

	user := &model.User{}
	if user, err = u.userSvc.GetUser(intCtx, uid); err != nil {
		zlog.With(intCtx).Warnw("UserSvc GetUser failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetUser OK", user)
}
