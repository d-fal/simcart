package repository

import (
	"context"
	"simcart/api/pb/cartpb"
	"simcart/domain/cart/entity"

	product_entity "simcart/domain/product/entity"
	"simcart/infrastructure/postgres"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cart struct {
	ctx      context.Context
	instance entity.CartAdapter
	tx       postgres.Tx
}
type CartOperations interface {
	getCurrentCart(cid uuid.UUID) error
	AddItem(context.Context, *cartpb.CartRequest) (*cartpb.CartResponse, error)
	getProduct(in *cartpb.CartRequest) (*product_entity.Product, error)
	ListCarts(in *cartpb.CartFilter) (*cartpb.CartRequests, error)
	AddCid(cid uuid.UUID) error
	RemoveItem(itemId uuid.UUID) (*emptypb.Empty, error)
	Checkout(cartUUID uuid.UUID) (*cartpb.CheckoutResponse, error)
}

func NewCart(ctx context.Context, owner uuid.UUID) (CartOperations, error) {
	c := new(cart)
	c.ctx = ctx

	tx, err := postgres.Storage.Transaction()

	if err != nil {
		return nil, status.Error(codes.Aborted, "db not initialized")
	}

	c.tx = tx
	c.instance = entity.NewCart()
	c.instance.SetOwner(owner)

	return c, nil
}
