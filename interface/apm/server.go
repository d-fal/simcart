package apm

import (
	"simcart/pkg/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func RegisterHealthServer() types.RegisterRPCInterface {
	return func(s *grpc.Server, cc *grpc.ClientConn, sm *runtime.ServeMux) error {
		healthgrpc.RegisterHealthServer(s, health.NewServer())
		return nil
	}
}
