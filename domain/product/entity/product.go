package entity

import (
	"encoding/json"
	"simcart/api/pb/commonpb"
	"simcart/pkg/model"

	"github.com/go-pg/pg/v10"
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
	SetSku(sku string) *Product
	Select() model.SelectOrInsertFunc
	Get() *Product
	Insert() model.InsertFunc
	SetCurrency(commonpb.Currency) *Product
	SetDetails(map[string]string) *Product
	SetCategory(cat commonpb.Category) *Product
	SetTitle(string) *Product
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

func (p *Product) SetSku(sku string) *Product {
	p.Sku = sku
	return p
}

func (p *Product) Select() model.SelectOrInsertFunc {
	return func(db *pg.DB) error {
		return db.Model(p).Where("sku = ?", p.Sku).Select()
	}
}

func (p *Product) Get() *Product {
	return p
}

func (p *Product) Insert() model.InsertFunc {

	return func(tx *pg.Tx) error {
		_, err := tx.Model(p).OnConflict("(sku) do update").Insert()

		return err
	}
}

func (p *Product) SetCurrency(cur commonpb.Currency) *Product {
	p.Currency = cur
	return p
}
func (p *Product) SetDetails(detail map[string]string) *Product {
	p.Descriptions = detail
	return p
}

func (p *Product) SetCategory(cat commonpb.Category) *Product {
	p.Category = cat
	return p
}

func (p *Product) SetTitle(title string) *Product {
	p.Title = title

	return p
}
