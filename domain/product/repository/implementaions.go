package repository

import (
	"context"
	"fmt"
	"simcart/api/pb/productpb"
	"simcart/domain/product/entity"
	"simcart/infrastructure/postgres"
	"simcart/infrastructure/search"
	"simcart/pkg/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *impl) Add(_ context.Context, in *productpb.Product) (*emptypb.Empty, error) {

	pHandler := entity.NewProduct()

	pHandler.SetPrice(in.Price).
		SetSku(in.Sku).
		SetCurrency(in.Currency).
		SetTitle(in.Product).
		SetCategory(in.Category).
		SetDetails(utils.ConvertValueToMap(in.Details))

	tx, err := postgres.Storage.Transaction()

	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot get db instance. %s", err.Error())
	}

	if err := tx.Begin(); err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot init db. %s", err.Error())
	}

	defer tx.Close()

	if err := pHandler.Insert()(tx.Get()); err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot insert in %s", err.Error())
	}
	// add to search db

	sClient, err := search.GetClient()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "search db is not available: %v", err)
	}

	serialized, _ := pHandler.Get().Marshal()
	doc := sClient.NewDocument(pHandler.Get().Sku, 0.9, pHandler.Get().Descriptions)
	doc.Set("date", pHandler.Get().CreatedAt.Unix()).
		Set("title", pHandler.Get().Title).
		Set("cat", pHandler.Get().Category.String()).
		Set("price", pHandler.Get().Price).
		Set("currency", pHandler.Get().Currency).
		Set("sku", pHandler.Get().Sku)

	// storing the whole object for seamless retrieval
	doc.SetPayload(serialized)

	if err := sClient.AddDocument(doc); err != nil {
		fmt.Printf("cannot create document %s", err.Error())

	}

	tx.Commit()

	return &emptypb.Empty{}, nil
}
