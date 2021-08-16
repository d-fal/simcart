package search_test

import (
	"context"
	"simcart/api/pb/commonpb"
	"simcart/api/pb/productpb"
	"simcart/api/pb/productpb/searchpb"
	"simcart/config"
	"simcart/infrastructure/postgres"
	search_db "simcart/infrastructure/search"
	"simcart/interface/product"
	"simcart/interface/product/search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	client := search_db.NewClient("redisearch:6379", "test")

	assert.NotNil(t, client)

	err := connectDb()

	assert.NoError(t, err)

	handler := search.NewSearchServerHandler(context.TODO())
	cases := []struct {
		name           string
		expected       interface{}
		success        bool
		productRequest *productpb.Product
		keyword        string
	}{
		{
			name:     "insert a new product and search it ",
			expected: commonpb.Category_Stationary,
			success:  true,
			keyword:  "pencil",
			productRequest: &productpb.Product{
				Product:  "Pencil",
				Category: commonpb.Category_Stationary,
				Sku:      "9000010001",
				Price:    1,
				Currency: commonpb.Currency_USD,
			},
		},
	}
	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			t.Run("insert", func(t *testing.T) {
				prodHandler := product.NewProductServerHandler(context.TODO())
				_, err := prodHandler.Add(context.TODO(), c.productRequest)

				assert.NoError(t, err)
			})

			t.Run("search", func(t *testing.T) {
				_, err := handler.SearchProduct(context.TODO(), &searchpb.Request{
					Keyword: c.keyword,
				})

				if c.success {
					assert.NoError(t, err)
				} else {
					assert.NotNil(t, err)
				}
			})
		})

	}

}

func connectDb() error {
	return postgres.Storage.Connect(config.Silent, &config.Postgres{
		Host:     "postgres",
		Port:     "5432",
		Database: "test_db",
		Password: "changeme",
		User:     "postgres",
	})
}
