package product

import (
	"context"
	"simcart/api/pb/productpb"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	handler := new(server)

	_, err := handler.Add(context.Background(), &productpb.Product{})

	assert.NotNil(t, err)
}
