package entity

import (
	"simcart/api/pb/cartpb"
	product_entity "simcart/domain/product/entity"
	"simcart/pkg/model"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type CartItem struct {
	model.Model
	UUID      uuid.UUID               `pg:",notnull,type:uuid default uuid_generate_v4()"`
	ProductId uint64                  `pg:",unique:idx_product_id_cart_id"`
	Product   *product_entity.Product `pg:"rel:has-one"`

	Discount float32 `pg:",use_zero"`
	Qty      uint64  `pg:",notnull"`

	CartId uint64 `pg:",notnull,unique:idx_product_id_cart_id"`
	Cart   *Cart  `pg:"rel:has-one"`
	model.Deleteables
}
type CartItems []*CartItem

type CartOperations interface {
	Add() model.InsertFunc
	SetProduct(product *product_entity.Product) *CartItem
	SetDiscount(discount float32) *CartItem
	SetQty(qty uint64) *CartItem
	SetCart(cart *Cart) *CartItem
	DropItem(itemId uuid.UUID) func(db *pg.DB) error
}

func NewCartItem() CartOperations {
	return new(CartItem)
}

func (c *CartItem) Add() model.InsertFunc {
	return func(tx *pg.Tx) error {
		if _, err := tx.Model(c).
			OnConflict("(product_id,cart_id) do update").
			Insert(); err != nil {
			return err
		}
		return nil
	}
}

func (c *CartItem) SetProduct(p *product_entity.Product) *CartItem {
	c.Product, c.ProductId = p, p.Id
	return c
}

func (c *CartItem) SetDiscount(discount float32) *CartItem {
	c.Discount = discount
	return c
}

func (c *CartItem) SetQty(qty uint64) *CartItem {
	c.Qty = qty
	return c
}

func (c *CartItem) SetCart(cart *Cart) *CartItem {
	c.Cart, c.CartId = cart, cart.Id
	return c
}

func (c *CartItem) DropItem(itemId uuid.UUID) func(db *pg.DB) error {
	return func(db *pg.DB) error {
		_, err := db.Model(c).Where("uuid = ?", itemId).Delete()
		return err
	}
}

// ToPb works as an adapter to prepare data to be used in grpc friendly way
func (cs CartItems) ToPb() []*cartpb.CartDetail {
	reuslt := []*cartpb.CartDetail{}
	for _, item := range cs {
		reuslt = append(reuslt, &cartpb.CartDetail{
			Qty:      item.Qty,
			Discount: item.Discount,
			ItemUUID: item.UUID.String(),
		})
	}

	return reuslt
}
