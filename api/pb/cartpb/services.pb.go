// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cartpb/services.proto

package cartpb

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_cartpb_services_proto protoreflect.FileDescriptor

var file_cartpb_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x69, 0x6d, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x15, 0x63, 0x61, 0x72, 0x74, 0x70,
	0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8a, 0x03, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x5d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1d, 0x2e, 0x73, 0x69,
	0x6d, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6d,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x64,
	0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x69, 0x6d,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6d, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x57, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x73,
	0x69, 0x6d, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x72, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x08, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x69, 0x6d, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x69, 0x6d, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x1a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x69, 0x6d, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x70,
	0x62, 0x3b, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cartpb_services_proto_goTypes = []interface{}{
	(*CartRequest)(nil),      // 0: simcart.api.cart.CartRequest
	(*CartFilter)(nil),       // 1: simcart.api.cart.CartFilter
	(*CartResponse)(nil),     // 2: simcart.api.cart.CartResponse
	(*CartRequests)(nil),     // 3: simcart.api.cart.CartRequests
	(*emptypb.Empty)(nil),    // 4: google.protobuf.Empty
	(*CheckoutResponse)(nil), // 5: simcart.api.cart.CheckoutResponse
}
var file_cartpb_services_proto_depIdxs = []int32{
	0, // 0: simcart.api.cart.Cart.Add:input_type -> simcart.api.cart.CartRequest
	1, // 1: simcart.api.cart.Cart.Get:input_type -> simcart.api.cart.CartFilter
	0, // 2: simcart.api.cart.Cart.Remove:input_type -> simcart.api.cart.CartRequest
	0, // 3: simcart.api.cart.Cart.Checkout:input_type -> simcart.api.cart.CartRequest
	2, // 4: simcart.api.cart.Cart.Add:output_type -> simcart.api.cart.CartResponse
	3, // 5: simcart.api.cart.Cart.Get:output_type -> simcart.api.cart.CartRequests
	4, // 6: simcart.api.cart.Cart.Remove:output_type -> google.protobuf.Empty
	5, // 7: simcart.api.cart.Cart.Checkout:output_type -> simcart.api.cart.CheckoutResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cartpb_services_proto_init() }
func file_cartpb_services_proto_init() {
	if File_cartpb_services_proto != nil {
		return
	}
	file_cartpb_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cartpb_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cartpb_services_proto_goTypes,
		DependencyIndexes: file_cartpb_services_proto_depIdxs,
	}.Build()
	File_cartpb_services_proto = out.File
	file_cartpb_services_proto_rawDesc = nil
	file_cartpb_services_proto_goTypes = nil
	file_cartpb_services_proto_depIdxs = nil
}
