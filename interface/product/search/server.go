package search

import (
	"context"
	"fmt"
	"simcart/api/pb/productpb/searchpb"
	"simcart/pkg/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	// it would be used for graceful shutdown purposes
	gctx context.Context
	// uinmplemented services
	searchpb.UnimplementedSearchServer
}

func NewSearchServer(ctx context.Context) types.RegisterRPCInterface {

	return func(s *grpc.Server, cc *grpc.ClientConn, sm *runtime.ServeMux) error {

		srv := new(server)
		srv.gctx = ctx

		searchpb.RegisterSearchServer(s, srv)

		if err := searchpb.RegisterSearchHandler(ctx, sm, cc); err != nil {
			return fmt.Errorf("Failed to connect to register gateway: %v\n", err)
		}

		return nil
	}
}
