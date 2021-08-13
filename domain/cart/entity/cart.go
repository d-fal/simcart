package cart

import (
	product_entity "simcart/domain/product/entity"
	"simcart/pkg/model"

	"github.com/go-pg/pg/v10"
)

type Cart struct {
	model.Model

	ProductId uint64
	Product   *product_entity.Product `pg:",rel:has-one"`

	Discount float32
	Qty      uint64

	Owner string // simply put, whoever makes the request

	model.Deleteables
}

type CartOperations interface {
	Add() model.InsertFunc
	SetProduct(product *product_entity.Product) *Cart
	SetDiscount(discount float32) *Cart
	SetQty(qty uint64) *Cart
}

func NewCart() CartOperations {
	return new(Cart)
}

func (c *Cart) Add() model.InsertFunc {
	return func(tx *pg.Tx) error {
		if _, err := tx.Model(c).Insert(); err != nil {
			return err
		}
		return nil
	}
}

func (c *Cart) SetProduct(p *product_entity.Product) *Cart {
	c.Product, c.ProductId = p, p.Id
	return c
}

func (c *Cart) SetDiscount(discount float32) *Cart {
	c.Discount = discount
	return c
}

func (c *Cart) SetQty(qty uint64) *Cart {
	c.Qty = qty
	return c
}
