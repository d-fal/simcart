package entity

import (
	"encoding/json"
	"simcart/api/pb/commonpb"
	"simcart/pkg/model"

	"github.com/google/uuid"
)

type Product struct {
	model.Model

	UUID uuid.UUID `pg:",notnull,type:uuid default uuid_generate_v4()"`

	Sku string `pg:",notnull,unique"`

	Category commonpb.Category `pg:",use_zero"`

	Price float64

	Currency commonpb.Currency

	Title string

	Descriptions map[string]string `pg:",hstore"`

	model.Deleteables
}

type ProductAdapter interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	SetPrice(price float64) *Product
}

func NewProduct() ProductAdapter {
	return new(Product)
}

func (p *Product) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *Product) SetPrice(price float64) *Product {
	p.Price = price
	return p
}
