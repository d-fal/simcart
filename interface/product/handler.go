package product

import (
	"context"
	"simcart/api/pb/productpb"
)

type serverHandler struct {
	// it would be used for graceful shutdown purposes
	ctx context.Context
	// uinmplemented services
	productpb.UnimplementedProductServiceServer
}

func NewProductServerHandler(ctx context.Context) *serverHandler {
	s := new(serverHandler)
	s.ctx = ctx
	return s
}
