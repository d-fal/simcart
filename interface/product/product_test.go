package product_test

import (
	"context"
	"simcart/api/pb/productpb"
	"simcart/interface/product"
	"testing"
)

func TestSearch(t *testing.T) {
	handler := product.NewProductServerHandler(context.TODO())

	handler.Add(context.Background(), &productpb.Product{})
}
