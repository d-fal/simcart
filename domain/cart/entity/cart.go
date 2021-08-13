package entity

import (
	"simcart/api/pb/cartpb"
	"simcart/pkg/model"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Cart struct {
	model.Model

	tableName struct{} `pg:"carts"`

	UUID  uuid.UUID `pg:",notnull,type:uuid default uuid_generate_v4()"`
	Owner uuid.UUID `pg:",unique:id_status_owner,notnull,type:uuid"` // simply put, whoever makes the request

	Status cartpb.CartStatus `pg:",unique:id_status_owner,use_zero"`

	CartItems []*CartItem `pg:"rel:has-many"`

	model.Deleteables
}
type CartAdapter interface {
	Select(cid uuid.UUID, tatus cartpb.CartStatus) model.SelectOrInsertFunc
	NewItem() CartOperations
	SetOwner(owner uuid.UUID)
	Get() *Cart
}

func NewCart() CartAdapter {
	return new(Cart)
}

func (c *Cart) Select(cid uuid.UUID, status cartpb.CartStatus) model.SelectOrInsertFunc {
	return func(db *pg.DB) error {
		_, err := db.Model(c).Where("owner = ? ", c.Owner).Where("status = ? ", status).OnConflict("do nothing").Returning("uuid").SelectOrInsert()
		return err
	}
}

func (c *Cart) NewItem() CartOperations {
	return newCartItem()
}

func (c *Cart) Get() *Cart {
	return c
}

func (c *Cart) SetOwner(owner uuid.UUID) {
	c.Owner = owner
}
