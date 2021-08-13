package utils

// func ToGRPC(val interface{}) *Value {
// 	switch val.(type) {
// 	case string:
// 		return &Value{
// 			Kind: &Value_StringValue{
// 				StringValue: val.(string),
// 			},
// 		}
// 	case int64:
// 		return &Value{
// 			Kind: &Value_IntValue{
// 				IntValue: val.(int64),
// 			},
// 		}
// 	case int32:
// 		return &Value{
// 			Kind: &Value_Int32Value{
// 				Int32Value: val.(int32),
// 			},
// 		}
// 	case int:
// 		return &Value{
// 			Kind: &Value_Int32Value{
// 				Int32Value: int32(val.(int)),
// 			},
// 		}
// 	case float32:
// 		return &Value{
// 			Kind: &Value_Float32Value{
// 				Float32Value: val.(float32),
// 			},
// 		}
// 	case float64:
// 		return &Value{
// 			Kind: &Value_Float64Value{
// 				Float64Value: val.(float64),
// 			},
// 		}
// 	case []interface{}:
// 		return &Value{
// 			Kind: &Value_ArrayValue{
// 				ArrayValue: &Array{Value: ToGRPCArray(val.([]interface{}))},
// 			},
// 		}
// 	default:
// 		return nil
// 	}
// }

// func ToGRPCArray(in []interface{}) []*Value {
// 	out := make([]*Value, len(in))

// 	for _, val := range in {
// 		out = append(out, ToGRPC(val))
// 	}
// 	return out
// }

// func ToGRPCMap(in map[string]interface{}) map[string]*Value {
// 	out := make(map[string]*Value)

// 	for key, val := range in {
// 		out[key] = ToGRPC(val)
// 	}
// 	return out
// }

// func (m *Value) Normalize() *interface{} {
// 	var val interface{}

// 	if x, ok := m.GetKind().(*Value_NullValue); ok {
// 		val = x.NullValue
// 	} else if x, ok := m.GetKind().(*Value_NumberValue); ok {
// 		val = x.NumberValue
// 	} else if x, ok := m.GetKind().(*Value_StringValue); ok {
// 		val = x.StringValue
// 	} else if x, ok := m.GetKind().(*Value_BoolValue); ok {
// 		val = x.BoolValue
// 	} else if x, ok := m.GetKind().(*Value_StructValue); ok {
// 		val = x.StructValue
// 	} else if x, ok := m.GetKind().(*Value_IntValue); ok {
// 		val = x.IntValue
// 	} else if x, ok := m.GetKind().(*Value_Int32Value); ok {
// 		val = x.Int32Value
// 	} else if x, ok := m.GetKind().(*Value_Float32Value); ok {
// 		val = x.Float32Value
// 	} else if x, ok := m.GetKind().(*Value_Float64Value); ok {
// 		val = x.Float64Value
// 	} else if x, ok := m.GetKind().(*Value_AnyValue); ok {
// 		val = x.AnyValue
// 	} else if x, ok := m.GetKind().(*Value_ArrayValue); ok {
// 		val = NormalizeArray(x.ArrayValue.GetValue())
// 	}

// 	return &val
// }

// func NormalizeArray(in []*Value) []interface{} {
// 	var out []interface{}

// 	for _, val := range in {
// 		if _val := val.Normalize(); _val != nil && *_val != nil {
// 			out = append(out, *_val)
// 		}
// 	}
// 	return out
// }

// func NormalizeMap(in map[string]*pb.Value) map[string]interface{} {
// 	out := make(map[string]interface{})

// 	for key, val := range in {
// 		out[key] = *val.Normalize()
// 	}
// 	return out
// }
