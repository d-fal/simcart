package repository

import (
	"context"
	"simcart/api/pb/productpb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductOperations interface {
	Add(context.Context, *productpb.Product) (*emptypb.Empty, error)
}

type impl struct {
}

func NewProductOperations() ProductOperations {
	return new(impl)
}
