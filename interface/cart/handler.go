package cart

import (
	"context"
	"simcart/api/pb/cartpb"
	cart_repository "simcart/domain/cart/repository"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	cartUUID, _ := uuid.Parse(in.CartUUID)
	owner, err := uuid.Parse(in.Owner)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field owner. %s", err.Error())
	}

	cart, err := cart_repository.NewCart(s.ctx, cartUUID, owner)

	if err != nil {
		return nil, err
	}
	return cart.AddItem(ctx, in)
}
