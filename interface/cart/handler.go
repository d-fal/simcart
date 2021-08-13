package cart

import (
	"context"
	"simcart/api/pb/cartpb"
	cart_repository "simcart/domain/cart/repository"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

	cartUUID, err := uuid.Parse(in.CartUUID)

	if err != nil && in.CartUUID != "" {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field cartUUID. %s", err.Error())
	}
	owner, err := uuid.Parse(in.Owner)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field owner. %s", err.Error())
	}

	cart, err := cart_repository.NewCart(s.ctx, owner)
	cart.AddCid(cartUUID)

	if err != nil {
		return nil, err
	}
	return cart.AddItem(ctx, in)
}

func (s *serverHandler) Get(ctx context.Context, in *cartpb.CartFilter) (*cartpb.CartRequests, error) {

	owner, err := uuid.Parse(in.Owner)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field owner. %s", err.Error())
	}

	cart, err := cart_repository.NewCart(s.ctx, owner)

	if err != nil {
		return nil, err
	}

	return cart.ListCarts(in)

}

func (s *serverHandler) Remove(ctx context.Context, in *cartpb.CartRequest) (*emptypb.Empty, error) {

	owner, err := uuid.Parse(in.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field owner. %s", err.Error())
	}

	cart, err := cart_repository.NewCart(s.ctx, owner)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot get cart instance. %s", err.Error())
	}

	itemId, err := uuid.Parse(in.ItemId)

	if err != nil {
		return nil, status.Errorf(codes.DataLoss, "cannot parse field item id. %s", err.Error())
	}

	return cart.RemoveItem(itemId)
}

func (s *serverHandler) Checkout(ctx context.Context, in *cartpb.CartRequest) (*cartpb.CheckoutResponse, error) {
	cartUUID, err := uuid.Parse(in.CartUUID)

	if err != nil {
		return nil, status.Errorf(codes.DataLoss, "cart uuid is not standard %s", err.Error())
	}
	owner, err := uuid.Parse(in.Owner)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "substandard uuid for the field owner. %s", err.Error())
	}

	cart, err := cart_repository.NewCart(s.ctx, owner)

	return cart.Checkout(cartUUID)

}
