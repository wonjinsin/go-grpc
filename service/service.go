package service

import (
	"context"
	"log"
	"os"
	"phantom/config"
	"phantom/repository"
	"phantom/util"

	"phantom/model"
)

var zlog *util.Logger

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Service ...
type Service struct {
	User UserService
}

// Init ...
func Init(conf *config.ViperConfig, repo *repository.Repository, redis *repository.RedisRepository) (*Service, error) {
	userSvc := NewUserService(repo.User)
	return &Service{
		User: userSvc,
	}, nil
}

// UserService ...
type UserService interface {
	GetUser(ctx context.Context, id string) (ruser *model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error)
	UpdateUser(ctx context.Context, uid string, user *model.User) (ruser *model.User, err error)
	DeleteUser(ctx context.Context, id string) (err error)
}
