package repository

import (
	"simcart/api/pb/productpb"
	"simcart/api/pb/productpb/searchpb"
	"simcart/clients/search"
	product_entity "simcart/domain/product/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Search(in *searchpb.Request, f func(*productpb.Product)) error {

	client, err := search.GetClient()

	if err != nil {
		return err
	}

	docs, _, err := client.Search(in.Keyword, "object", "title", "weight", "sku", "cat", "color")

	for _, doc := range docs {

		product := new(product_entity.Product)
		if err := product.Unmarshal(doc.Payload); err != nil {
			return status.Errorf(codes.DataLoss, "payload not set")
		}

		f(&productpb.Product{
			Product:  product.Title,
			Sku:      product.Sku,
			Category: product.Category,
			Price:    float64(product.Price),
			Currency: product.Currency,
		})

	}
	return err
}
