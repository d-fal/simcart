package types

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type RegisterRPCInterface func(s *grpc.Server, cc *grpc.ClientConn, sm *runtime.ServeMux) error

type SelectedInterfaces func(suggestedHandlers ...RegisterRPCInterface) error
