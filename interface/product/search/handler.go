package search

import (
	"context"
	"simcart/api/pb/productpb"
	"simcart/api/pb/productpb/searchpb"
	search_repository "simcart/domain/product/search/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverHandler struct {
	// it would be used for graceful shutdown purposes
	ctx context.Context
	// uinmplemented services
	searchpb.UnimplementedSearchServer
}

func NewSearchServerHandler(ctx context.Context) *serverHandler {
	s := new(serverHandler)
	s.ctx = ctx
	return s
}

func (s *serverHandler) SearchProduct(ctx context.Context, in *searchpb.Request) (*searchpb.Response, error) {
	searchResults := []*productpb.Product{}
	callBack := func(r *productpb.Product) {
		searchResults = append(searchResults, r)

	}

	if err := search_repository.Search(in, callBack); err != nil {
		return nil, status.Errorf(codes.NotFound, "pattern not found")
	}

	return &searchpb.Response{Result: searchResults}, nil
}
