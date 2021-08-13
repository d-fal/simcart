package entity

import (
	product_entity "simcart/domain/product/entity"
	"simcart/pkg/model"

	"github.com/go-pg/pg/v10"
)

type CartItem struct {
	model.Model

	ProductId uint64
	Product   *product_entity.Product `pg:"rel:has-one"`

	Discount float32
	Qty      uint64

	CartId uint64 `pg:",notnull"`
	Cart   *Cart  `pg:"rel:has-one"`
	model.Deleteables
}

type CartOperations interface {
	Add() model.InsertFunc
	SetProduct(product *product_entity.Product) *CartItem
	SetDiscount(discount float32) *CartItem
	SetQty(qty uint64) *CartItem
}

func newCartItem() CartOperations {
	return new(CartItem)
}

func (c *CartItem) Add() model.InsertFunc {
	return func(tx *pg.Tx) error {
		if _, err := tx.Model(c).Insert(); err != nil {
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
