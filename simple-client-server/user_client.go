package simple_client_server

import (
	userpb "github.com/heekng/go_grpc/protos/v1/user"
	"google.golang.org/grpc"
	"sync"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost, grpc.WithInsecure(), grpc.WithBlock())

		cli = userpb.NewUserClient(conn)
	})

	return cli
}
