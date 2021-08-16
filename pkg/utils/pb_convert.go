package utils

import "google.golang.org/protobuf/types/known/structpb"

func ConvertValueToMap(in *structpb.Struct) map[string]string {

	res := make(map[string]string)

	if in == nil {
		return res
	}
	for k, v := range in.Fields {
		res[k] = v.GetStringValue()
	}

	return res
}

func ConvertMapToValue(in map[string]string) *structpb.Struct {

	fields := make(map[string]*structpb.Value)

	for key, val := range in {
		fields[key] = &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: val,
			},
		}
	}

	return &structpb.Struct{
		Fields: fields,
	}
}
