package utils_test

import (
	"simcart/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestConversionToMap(t *testing.T) {

	sample := []*structpb.Struct{
		{
			Fields: map[string]*structpb.Value{
				"v1": {
					Kind: &structpb.Value_StringValue{
						StringValue: "v1",
					},
				},
			},
		},
		{
			Fields: map[string]*structpb.Value{
				"v3": {
					Kind: &structpb.Value_StringValue{
						StringValue: "v3",
					},
				},
			},
		},
		{
			Fields: map[string]*structpb.Value{
				"new key": {
					Kind: &structpb.Value_StringValue{
						StringValue: "new key",
					},
				},
			},
		},
	}

	for _, s := range sample {

		for k, v := range s.Fields {
			assert.Equal(t, k, v.GetStringValue())
		}

		for k, v := range utils.ConvertValueToMap(s) {
			assert.Equal(t, k, v)
		}

	}

}

func TestConversionToStruct(t *testing.T) {
	cases := map[string]string{
		"name1": "name1",
		"name2": "name2",
		"30000": "30000",
	}

	t.Run("check equality of converted map[string]string", func(t *testing.T) {
		convertedMap := utils.ConvertMapToValue(cases)
		for k, v := range convertedMap.Fields {
			assert.Equal(t, k, v.GetStringValue())
		}

	})

}
