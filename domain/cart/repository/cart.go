package repository

import (
	"context"
	"simcart/api/pb/cartpb"
	"simcart/domain/cart/entity"

	product_entity "simcart/domain/product/entity"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *cart) AddItem(ctx context.Context, in *cartpb.CartRequest) (*cartpb.CartResponse, error) {
	c.tx.Begin()

	defer c.tx.Close()

	product, err := c.getProduct(in)

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "product %s cannot be fetched from db", in.Sku)
	}

	newItem := c.instance.NewItem().
		SetDiscount(0).SetQty(uint64(in.Qty)).
		SetCart(c.instance.Get()).SetProduct(product)

	if err := newItem.Add()(c.tx.Get()); err != nil {
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

func (c *cart) getProduct(in *cartpb.CartRequest) (*product_entity.Product, error) {
	product := product_entity.NewProduct()

	if err := product.SetSku(in.Sku).Select()(c.tx.Db()); err != nil {
		return nil, err
	}

	return product.Get(), nil
}

func (c *cart) ListCarts(in *cartpb.CartFilter) (*cartpb.CartRequests, error) {

	owner, err := uuid.Parse(in.Owner)

	if err != nil {
		return nil, status.Errorf(codes.DataLoss, "cannot parse field owner. %s", err.Error())
	}

	list, err := entity.Filter(owner, in.Status...)(c.tx.Db())

	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot list cart requests : %s", err.Error())
	}

	return &cartpb.CartRequests{
		Responses: list.ToPb(),
	}, nil
}

func (c *cart) RemoveItem(itemId uuid.UUID) (*emptypb.Empty, error) {

	newCartItem := entity.NewCartItem()

	return &emptypb.Empty{}, newCartItem.DropItem(itemId)(c.tx.Db())
}

func (c *cart) AddCid(cid uuid.UUID) error {
	return c.getCurrentCart(cid)
}

func (c *cart) Checkout(cartUUID uuid.UUID) (*cartpb.CheckoutResponse, error) {
	if err := c.instance.Update(cartUUID, cartpb.CartStatus_CheckedOut)(c.tx.Db()); err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot update cart %s", err.Error())
	}

	return &cartpb.CheckoutResponse{
		CartUUID:   cartUUID.String(),
		PaymentUrl: "https://bank.example.com/" + cartUUID.String(),
	}, nil
}
