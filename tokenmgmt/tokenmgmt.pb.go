// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: tokenmgmt/tokenmgmt.proto

package go_tokenmgmt_grpc

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

type TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *TokenData) Reset() {
	*x = TokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenData) ProtoMessage() {}

func (x *TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenData.ProtoReflect.Descriptor instead.
func (*TokenData) Descriptor() ([]byte, []int) {
	return file_tokenmgmt_tokenmgmt_proto_rawDescGZIP(), []int{0}
}

func (x *TokenData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TokenData) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *TokenData) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type WriteData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Low  uint64 `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	Mid  uint64 `protobuf:"varint,5,opt,name=mid,proto3" json:"mid,omitempty"`
	High uint64 `protobuf:"varint,6,opt,name=high,proto3" json:"high,omitempty"`
}

func (x *WriteData) Reset() {
	*x = WriteData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteData) ProtoMessage() {}

func (x *WriteData) ProtoReflect() protoreflect.Message {
	mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteData.ProtoReflect.Descriptor instead.
func (*WriteData) Descriptor() ([]byte, []int) {
	return file_tokenmgmt_tokenmgmt_proto_rawDescGZIP(), []int{1}
}

func (x *WriteData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WriteData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteData) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *WriteData) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WriteData) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WriteData) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

type SuccessStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SuccessStatus) Reset() {
	*x = SuccessStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessStatus) ProtoMessage() {}

func (x *SuccessStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessStatus.ProtoReflect.Descriptor instead.
func (*SuccessStatus) Descriptor() ([]byte, []int) {
	return file_tokenmgmt_tokenmgmt_proto_rawDescGZIP(), []int{2}
}

func (x *SuccessStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ResultRead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Finalval uint64 `protobuf:"varint,1,opt,name=finalval,proto3" json:"finalval,omitempty"`
}

func (x *ResultRead) Reset() {
	*x = ResultRead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRead) ProtoMessage() {}

func (x *ResultRead) ProtoReflect() protoreflect.Message {
	mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRead.ProtoReflect.Descriptor instead.
func (*ResultRead) Descriptor() ([]byte, []int) {
	return file_tokenmgmt_tokenmgmt_proto_rawDescGZIP(), []int{3}
}

func (x *ResultRead) GetFinalval() uint64 {
	if x != nil {
		return x.Finalval
	}
	return 0
}

type ResultWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partialval uint64 `protobuf:"varint,1,opt,name=partialval,proto3" json:"partialval,omitempty"`
}

func (x *ResultWrite) Reset() {
	*x = ResultWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultWrite) ProtoMessage() {}

func (x *ResultWrite) ProtoReflect() protoreflect.Message {
	mi := &file_tokenmgmt_tokenmgmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultWrite.ProtoReflect.Descriptor instead.
func (*ResultWrite) Descriptor() ([]byte, []int) {
	return file_tokenmgmt_tokenmgmt_proto_rawDescGZIP(), []int{4}
}

func (x *ResultWrite) GetPartialval() uint64 {
	if x != nil {
		return x.Partialval
	}
	return 0
}

var File_tokenmgmt_tokenmgmt_proto protoreflect.FileDescriptor

var file_tokenmgmt_tokenmgmt_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x43, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x7b, 0x0a, 0x09, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c,
	0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x22, 0x21, 0x0a, 0x0d, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x28, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x76, 0x61, 0x6c, 0x22, 0x2d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x76, 0x61, 0x6c, 0x32, 0xf7, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x15, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x61, 0x64, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x05, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b,
	0x67, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tokenmgmt_tokenmgmt_proto_rawDescOnce sync.Once
	file_tokenmgmt_tokenmgmt_proto_rawDescData = file_tokenmgmt_tokenmgmt_proto_rawDesc
)

func file_tokenmgmt_tokenmgmt_proto_rawDescGZIP() []byte {
	file_tokenmgmt_tokenmgmt_proto_rawDescOnce.Do(func() {
		file_tokenmgmt_tokenmgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_tokenmgmt_tokenmgmt_proto_rawDescData)
	})
	return file_tokenmgmt_tokenmgmt_proto_rawDescData
}

var file_tokenmgmt_tokenmgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tokenmgmt_tokenmgmt_proto_goTypes = []interface{}{
	(*TokenData)(nil),     // 0: tokenmgmt.TokenData
	(*WriteData)(nil),     // 1: tokenmgmt.WriteData
	(*SuccessStatus)(nil), // 2: tokenmgmt.SuccessStatus
	(*ResultRead)(nil),    // 3: tokenmgmt.ResultRead
	(*ResultWrite)(nil),   // 4: tokenmgmt.ResultWrite
}
var file_tokenmgmt_tokenmgmt_proto_depIdxs = []int32{
	0, // 0: tokenmgmt.TokenManagement.Create:input_type -> tokenmgmt.TokenData
	0, // 1: tokenmgmt.TokenManagement.Read:input_type -> tokenmgmt.TokenData
	1, // 2: tokenmgmt.TokenManagement.Write:input_type -> tokenmgmt.WriteData
	0, // 3: tokenmgmt.TokenManagement.Drop:input_type -> tokenmgmt.TokenData
	2, // 4: tokenmgmt.TokenManagement.Create:output_type -> tokenmgmt.SuccessStatus
	3, // 5: tokenmgmt.TokenManagement.Read:output_type -> tokenmgmt.ResultRead
	4, // 6: tokenmgmt.TokenManagement.Write:output_type -> tokenmgmt.ResultWrite
	2, // 7: tokenmgmt.TokenManagement.Drop:output_type -> tokenmgmt.SuccessStatus
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tokenmgmt_tokenmgmt_proto_init() }
func file_tokenmgmt_tokenmgmt_proto_init() {
	if File_tokenmgmt_tokenmgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tokenmgmt_tokenmgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenmgmt_tokenmgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenmgmt_tokenmgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenmgmt_tokenmgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRead); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenmgmt_tokenmgmt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultWrite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tokenmgmt_tokenmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tokenmgmt_tokenmgmt_proto_goTypes,
		DependencyIndexes: file_tokenmgmt_tokenmgmt_proto_depIdxs,
		MessageInfos:      file_tokenmgmt_tokenmgmt_proto_msgTypes,
	}.Build()
	File_tokenmgmt_tokenmgmt_proto = out.File
	file_tokenmgmt_tokenmgmt_proto_rawDesc = nil
	file_tokenmgmt_tokenmgmt_proto_goTypes = nil
	file_tokenmgmt_tokenmgmt_proto_depIdxs = nil
}