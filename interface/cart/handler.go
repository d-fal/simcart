package cart

import (
	"context"
	"simcart/api/pb/cartpb"
)

type serverHandler struct {
	// it would be used for graceful shutdown purposes
	ctx context.Context
	// uinmplemented services
	cartpb.UnimplementedCartServer
}

func NewCartServerHandler(ctx context.Context) *serverHandler {
	s := new(serverHandler)
	s.ctx = ctx
	return s
}

func (s *serverHandler) Add(ctx context.Context, in *cartpb.CartRequest) (*cartpb.CartResponse, error) {

	return nil, nil
}
