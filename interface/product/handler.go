package product

import (
	"context"
	"simcart/api/pb/productpb"
	"simcart/domain/product/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) Add(ctx context.Context, in *productpb.Product) (*emptypb.Empty, error) {

	npo := repository.NewProductOperations()

	return npo.Add(ctx, in)
}
