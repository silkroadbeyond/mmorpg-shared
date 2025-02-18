package grpc

import (
	"github.com/silkroadbeyond/mmorpg-shared/logging"
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryInterceptor,
		),
	)
}

func NewClient(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}
