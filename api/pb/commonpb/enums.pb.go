// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: commonpb/enums.proto

package commonpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type Category int32

const (
	Category_Appliance  Category = 0
	Category_Stationary Category = 1
	Category_FMCG       Category = 2
	Category_Apparel    Category = 3
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "Appliance",
		1: "Stationary",
		2: "FMCG",
		3: "Apparel",
	}
	Category_value = map[string]int32{
		"Appliance":  0,
		"Stationary": 1,
		"FMCG":       2,
		"Apparel":    3,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_commonpb_enums_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_commonpb_enums_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_commonpb_enums_proto_rawDescGZIP(), []int{0}
}

type Currency int32

const (
	Currency_USD Currency = 0
	Currency_EUR Currency = 1
	Currency_GBP Currency = 2
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "USD",
		1: "EUR",
		2: "GBP",
	}
	Currency_value = map[string]int32{
		"USD": 0,
		"EUR": 1,
		"GBP": 2,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_commonpb_enums_proto_enumTypes[1].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_commonpb_enums_proto_enumTypes[1]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_commonpb_enums_proto_rawDescGZIP(), []int{1}
}

var File_commonpb_enums_proto protoreflect.FileDescriptor

var file_commonpb_enums_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x69, 0x6d, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x40, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4d, 0x43, 0x47, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x10, 0x03, 0x2a, 0x2b, 0x0a, 0x08,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x55, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x42,
	0x50, 0x10, 0x02, 0x22, 0x04, 0x08, 0x03, 0x10, 0x64, 0x42, 0x22, 0x5a, 0x20, 0x73, 0x69, 0x6d,
	0x63, 0x61, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonpb_enums_proto_rawDescOnce sync.Once
	file_commonpb_enums_proto_rawDescData = file_commonpb_enums_proto_rawDesc
)

func file_commonpb_enums_proto_rawDescGZIP() []byte {
	file_commonpb_enums_proto_rawDescOnce.Do(func() {
		file_commonpb_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonpb_enums_proto_rawDescData)
	})
	return file_commonpb_enums_proto_rawDescData
}

var file_commonpb_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commonpb_enums_proto_goTypes = []interface{}{
	(Category)(0), // 0: simcart.api.common.Category
	(Currency)(0), // 1: simcart.api.common.Currency
}
var file_commonpb_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonpb_enums_proto_init() }
func file_commonpb_enums_proto_init() {
	if File_commonpb_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonpb_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonpb_enums_proto_goTypes,
		DependencyIndexes: file_commonpb_enums_proto_depIdxs,
		EnumInfos:         file_commonpb_enums_proto_enumTypes,
	}.Build()
	File_commonpb_enums_proto = out.File
	file_commonpb_enums_proto_rawDesc = nil
	file_commonpb_enums_proto_goTypes = nil
	file_commonpb_enums_proto_depIdxs = nil
}
