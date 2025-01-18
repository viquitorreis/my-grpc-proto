// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.12
// source: proto/bank/type/exchange.proto

package bank

import (
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

type ExchangeRateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,proto3" json:"from_currency,omitempty"`
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,proto3" json:"to_currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeRateRequest) Reset() {
	*x = ExchangeRateRequest{}
	mi := &file_proto_bank_type_exchange_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateRequest) ProtoMessage() {}

func (x *ExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_exchange_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRateRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

type ExchangeRateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,proto3" json:"from_currency,omitempty"`
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,proto3" json:"to_currency,omitempty"`
	Rate          float64                `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Timestamp     string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeRateResponse) Reset() {
	*x = ExchangeRateResponse{}
	mi := &file_proto_bank_type_exchange_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateResponse) ProtoMessage() {}

func (x *ExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_exchange_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*ExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRateResponse) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ExchangeRateResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *ExchangeRateResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_bank_type_exchange_proto protoreflect.FileDescriptor

var file_proto_bank_type_exchange_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0x5d, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x71, 0x75, 0x69, 0x74, 0x6f, 0x72, 0x72,
	0x65, 0x69, 0x73, 0x2f, 0x6d, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bank_type_exchange_proto_rawDescOnce sync.Once
	file_proto_bank_type_exchange_proto_rawDescData = file_proto_bank_type_exchange_proto_rawDesc
)

func file_proto_bank_type_exchange_proto_rawDescGZIP() []byte {
	file_proto_bank_type_exchange_proto_rawDescOnce.Do(func() {
		file_proto_bank_type_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bank_type_exchange_proto_rawDescData)
	})
	return file_proto_bank_type_exchange_proto_rawDescData
}

var file_proto_bank_type_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_bank_type_exchange_proto_goTypes = []any{
	(*ExchangeRateRequest)(nil),  // 0: bank.ExchangeRateRequest
	(*ExchangeRateResponse)(nil), // 1: bank.ExchangeRateResponse
}
var file_proto_bank_type_exchange_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_bank_type_exchange_proto_init() }
func file_proto_bank_type_exchange_proto_init() {
	if File_proto_bank_type_exchange_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bank_type_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bank_type_exchange_proto_goTypes,
		DependencyIndexes: file_proto_bank_type_exchange_proto_depIdxs,
		MessageInfos:      file_proto_bank_type_exchange_proto_msgTypes,
	}.Build()
	File_proto_bank_type_exchange_proto = out.File
	file_proto_bank_type_exchange_proto_rawDesc = nil
	file_proto_bank_type_exchange_proto_goTypes = nil
	file_proto_bank_type_exchange_proto_depIdxs = nil
}
