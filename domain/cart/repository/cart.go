package repository

import (
	"context"
	"simcart/api/pb/cartpb"
	"simcart/domain/cart/entity"

	"simcart/clients/postgres"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartOperations interface {
	getCurrentCart(cid uuid.UUID) error
	AddItem(context.Context, *cartpb.CartRequest) (*cartpb.CartResponse, error)
}

type cart struct {
	ctx      context.Context
	instance entity.CartAdapter
	tx       postgres.Tx
}

func NewCart(ctx context.Context, cid uuid.UUID, owner uuid.UUID) (CartOperations, error) {
	c := new(cart)
	c.ctx = ctx

	tx, err := postgres.Storage.Transaction()

	if err != nil {
		return nil, status.Error(codes.Aborted, "db not initialized")
	}

	c.tx = tx
	c.instance = entity.NewCart()
	c.instance.SetOwner(owner)

	return c, c.getCurrentCart(cid)
}

func (c *cart) AddItem(ctx context.Context, in *cartpb.CartRequest) (*cartpb.CartResponse, error) {
	c.tx.Begin()

	defer c.tx.Close()

	if err := c.instance.NewItem().Add()(c.tx.Get()); err != nil {
		return nil, err
	}

	c.tx.Commit()

	return &cartpb.CartResponse{
		CartUUID: c.instance.Get().UUID.String(),
	}, nil
}

func (c *cart) getCurrentCart(cid uuid.UUID) error {

	if err := c.instance.Select(cid, cartpb.CartStatus_New)(c.tx.Db()); err != nil {
		return status.Errorf(codes.Aborted, "cannot select or insert | %s", err.Error())
	}
	return nil
}
