package logging

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var Logger *zap.Logger

func Init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	Logger.Info("gRPC call", zap.String("method", info.FullMethod))
	return handler(ctx, req)
}