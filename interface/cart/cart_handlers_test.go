package cart_test

import (
	"context"
	"testing"

	"simcart/api/pb/cartpb"
	"simcart/api/pb/commonpb"
	postgres "simcart/clients/postgres"
	"simcart/config"
	"simcart/interface/cart"

	"github.com/stretchr/testify/assert"
)

func connectDb() error {
	return postgres.Storage.Connect(config.Verbose, &config.Postgres{
		Host:     "postgres",
		Port:     "5432",
		Database: "test_db",
		Password: "changeme",
		User:     "postgres",
	})
}

func TestAddToCart(t *testing.T) {

	t.Run("testing db connection", func(t *testing.T) {
		var (
			owner   string
			cid     string
			err     error
			request *cartpb.CartRequest
		)

		owner = "b4a76de0-fc20-11eb-9a03-0242ac130003"

		err = connectDb()

		request = &cartpb.CartRequest{
			ProductUUID: "b4a76ef8-fc20-11eb-9a03-0242ac130003",
			Sku:         "1000000",
			Bid:         100,
			CartUUID:    cid,
			Owner:       owner,
			Cuurency:    commonpb.Currency_EUR,
		}

		assert.NotNil(t, err)

		t.Run("create a new cart item", func(t *testing.T) {

			handler := cart.NewCartServerHandler(context.TODO())
			response, err := handler.Add(context.TODO(), request)

			assert.NoError(t, err)
			assert.NotNil(t, response)

		})
	})
	connectDb()
}
