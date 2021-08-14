package cart_test

import (
	"context"
	"fmt"
	"testing"

	"simcart/api/pb/cartpb"
	"simcart/api/pb/commonpb"
	"simcart/config"
	postgres "simcart/infrastructure/postgres"
	"simcart/interface/cart"

	"github.com/logrusorgru/aurora"
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

var (
	owner = "b4a76de0-fc20-11eb-9a03-0242ac130003"
	cid   = "b4a76ef8-fc20-11eb-9a03-0242ac130003"
	err   error
)

func TestAddToCart(t *testing.T) {

	t.Run("testing db connection", func(t *testing.T) {

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
	t.Run("list carts", func(t *testing.T) {
		err := connectDb()

		assert.NoError(t, err)
		handler := cart.NewCartServerHandler(context.TODO())

		response, err := handler.Get(context.TODO(), &cartpb.CartFilter{
			Owner: owner,
			Status: []cartpb.CartStatus{
				cartpb.CartStatus_New,
			},
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("list carts", func(t *testing.T) {
		err := connectDb()

		assert.NoError(t, err)
		handler := cart.NewCartServerHandler(context.TODO())

		response, err := handler.Get(context.TODO(), &cartpb.CartFilter{
			Owner: owner,
			Status: []cartpb.CartStatus{
				cartpb.CartStatus_New,
			},
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		t.Run("remove cart ", func(t *testing.T) {

			for _, resp := range response.Responses {
				for _, detail := range resp.Details {
					_, err := handler.Remove(context.TODO(), &cartpb.CartRequest{
						Owner:  owner,
						ItemId: detail.ItemUUID,
					})

					assert.NoError(t, err)
				}
			}
		})

	})
}

func TestCheckout(t *testing.T) {
	err := connectDb()

	assert.NoError(t, err)
	handler := cart.NewCartServerHandler(context.TODO())

	t.Run("create a request", func(t *testing.T) {
		cartResponse, err := handler.Add(context.TODO(), &cartpb.CartRequest{
			Sku:      "100200300",
			Bid:      20,
			CartUUID: cid,
			Owner:    owner,
			Cuurency: commonpb.Currency_EUR,
			Qty:      3,
		})

		fmt.Println("order created ", err)

		assert.NoError(t, err)
		t.Run("checking out the request", func(t *testing.T) {
			_, err = handler.Checkout(context.TODO(), &cartpb.CartRequest{
				Owner:    owner,
				CartUUID: cartResponse.CartUUID,
			})

			assert.NoError(t, err)
			fmt.Println("order created ", aurora.Yellow(err))
		})
	})

}
