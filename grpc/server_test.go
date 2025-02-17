package grpc_test

import (
	"testing"

	"github.com/silkroadbeyond/mmorpg-shared/grpc"
	// "google.golang.org/grpc"
)

func TestNewServer(t *testing.T) {
	server := grpc.NewServer()
	if server == nil {
		t.Fatal("Failed to create gRPC server")
	}
}
