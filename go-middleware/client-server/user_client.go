package clientserver

import (
	"sync"

	"google.golang.org/grpc"

	"go-grpc-v1/protos/user"
)

var (
	once sync.Once
	cli  user.UserClient
)

// GetUserClient ...
func GetUserClient(serviceHost string) user.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)

		cli = user.NewUserClient(conn)
	})

	return cli
}
