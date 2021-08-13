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

	CartItems CartItems `pg:"rel:has-many"`

	model.Deleteables
}

type Carts []*Cart

type CartAdapter interface {
	Select(cid uuid.UUID, tatus cartpb.CartStatus) model.SelectOrInsertFunc
	NewItem() CartOperations
	SetOwner(owner uuid.UUID)
	Get() *Cart
	Update(cartUUID uuid.UUID, status cartpb.CartStatus) func(db *pg.DB) error
}

func NewCart() CartAdapter {
	return new(Cart)
}

func (c *Cart) Select(cid uuid.UUID, status cartpb.CartStatus) model.SelectOrInsertFunc {
	return func(db *pg.DB) error {
		_, err := db.Model(c).Where("owner = ? ", c.Owner).Where("status = ? ", status).
			OnConflict("do nothing").Returning("uuid").Returning("id").
			SelectOrInsert()
		return err
	}
}

func (c *Cart) NewItem() CartOperations {
	return NewCartItem()
}

func (c *Cart) Get() *Cart {
	return c
}

func (c *Cart) SetOwner(owner uuid.UUID) {
	c.Owner = owner
}

func (c *Cart) Update(cartUUID uuid.UUID, status cartpb.CartStatus) func(db *pg.DB) error {
	return func(db *pg.DB) error {
		_, err := db.Model(c).Where("uuid = ? ", cartUUID).Set("status = ? ", status).Update()
		return err
	}
}

func Filter(owner uuid.UUID, status ...cartpb.CartStatus) func(db *pg.DB) (Carts, error) {
	return func(db *pg.DB) (Carts, error) {
		var carts []*Cart

		query := db.Model(&carts).Where("owner = ?", owner)
		for _, s := range status {
			query = query.Where("status = ?", s)
		}

		if err := query.Relation("CartItems").Select(); err != nil {
			return nil, err
		}

		return carts, nil
	}
}

func (carts Carts) ToPb() []*cartpb.CartResponse {
	reposnse := []*cartpb.CartResponse{}

	for _, cart := range carts {
		reposnse = append(reposnse, &cartpb.CartResponse{
			CartUUID: cart.UUID.String(),
			Status:   cart.Status,
			Details:  cart.CartItems.ToPb(),
		})
	}

	return reposnse
}
