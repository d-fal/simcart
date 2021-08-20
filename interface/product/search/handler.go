package search

import (
	"context"
	"simcart/api/pb/productpb"
	"simcart/api/pb/productpb/searchpb"
	search_repository "simcart/domain/product/search/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) SearchProduct(ctx context.Context, in *searchpb.Request) (*searchpb.Response, error) {
	searchResults := []*productpb.Product{}
	callBack := func(r *productpb.Product) {
		searchResults = append(searchResults, r)

	}

	if err := search_repository.Search(in, callBack); err != nil {
		return nil, status.Errorf(codes.NotFound, "pattern not found")
	}

	return &searchpb.Response{Result: searchResults}, nil
}
