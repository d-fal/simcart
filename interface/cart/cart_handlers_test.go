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
	return postgres.Storage.Connect(config.Silent, &config.Postgres{
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
			owner string
			cid   string
			err   error
		)

		owner = "b4a76de0-fc20-11eb-9a03-0242ac130003"
		cid = "b4a76ef8-fc20-11eb-9a03-0242ac130003"
		err = connectDb()

		assert.NoError(t, err)

		requests := []struct {
			Case    string
			request *cartpb.CartRequest
			success bool
		}{
			{
				Case: "all the inputs are provided",
				request: &cartpb.CartRequest{
					Sku:      "100200200",
					Bid:      100,
					CartUUID: cid,
					Owner:    owner,
					Cuurency: commonpb.Currency_EUR,
					Qty:      4,
				},
				success: true,
			},
			{
				Case: "qty is missing",
				request: &cartpb.CartRequest{
					Sku:      "100200200",
					Bid:      100,
					CartUUID: cid,
					Owner:    owner,
					Cuurency: commonpb.Currency_EUR,
					// Qty:         1,
				},
				success: false,
			},
			{
				Case: "owner is missing",
				request: &cartpb.CartRequest{
					Sku:      "100200200",
					Bid:      100,
					CartUUID: cid,
					// Owner:       owner,
					Cuurency: commonpb.Currency_EUR,
					Qty:      1,
				},
				success: false,
			},
			{
				Case: "sku does not exist",
				request: &cartpb.CartRequest{
					Sku:      "1234",
					Bid:      100,
					CartUUID: cid,
					Owner:    owner,
					Cuurency: commonpb.Currency_EUR,
					Qty:      1,
				},
				success: false,
			},
			{
				Case: "cart  cid is not standard",
				request: &cartpb.CartRequest{
					Sku:      "100200200",
					Bid:      100,
					CartUUID: "0",
					Owner:    owner,
					Cuurency: commonpb.Currency_EUR,
					Qty:      1,
				},
				success: false,
			},
		}

		for _, request := range requests {
			t.Run(request.Case, func(t *testing.T) {
				handler := cart.NewCartServerHandler(context.TODO())

				response, err := handler.Add(context.TODO(), request.request)

				if request.success {
					assert.NoError(t, err)
					assert.NotEmpty(t, response)
				} else {
					assert.Error(t, err)
				}

			})
		}

	})

}

func TestListCart(t *testing.T) {

}
