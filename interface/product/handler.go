package product

import (
	"context"
	"simcart/api/pb/productpb"
	"simcart/domain/product/repository"

	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *serverHandler) Add(ctx context.Context, in *productpb.Product) (*emptypb.Empty, error) {

	npo := repository.NewProductOperations()

	return npo.Add(ctx, in)
}
