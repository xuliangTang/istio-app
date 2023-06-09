// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: UserKind.proto

package v1

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

type KindLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *UserLoginModel `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *KindLoginRequest) Reset() {
	*x = KindLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserKind_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KindLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KindLoginRequest) ProtoMessage() {}

func (x *KindLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserKind_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KindLoginRequest.ProtoReflect.Descriptor instead.
func (*KindLoginRequest) Descriptor() ([]byte, []int) {
	return file_UserKind_proto_rawDescGZIP(), []int{0}
}

func (x *KindLoginRequest) GetSpec() *UserLoginModel {
	if x != nil {
		return x.Spec
	}
	return nil
}

type KindLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *UserModel `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *KindLoginResponse) Reset() {
	*x = KindLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserKind_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KindLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KindLoginResponse) ProtoMessage() {}

func (x *KindLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserKind_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KindLoginResponse.ProtoReflect.Descriptor instead.
func (*KindLoginResponse) Descriptor() ([]byte, []int) {
	return file_UserKind_proto_rawDescGZIP(), []int{1}
}

func (x *KindLoginResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *KindLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *KindLoginResponse) GetData() *UserModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type KindRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *UserRegisterModel `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *KindRegisterRequest) Reset() {
	*x = KindRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserKind_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KindRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KindRegisterRequest) ProtoMessage() {}

func (x *KindRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_UserKind_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KindRegisterRequest.ProtoReflect.Descriptor instead.
func (*KindRegisterRequest) Descriptor() ([]byte, []int) {
	return file_UserKind_proto_rawDescGZIP(), []int{2}
}

func (x *KindRegisterRequest) GetSpec() *UserRegisterModel {
	if x != nil {
		return x.Spec
	}
	return nil
}

type KindRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *UserModel `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *KindRegisterResponse) Reset() {
	*x = KindRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserKind_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KindRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KindRegisterResponse) ProtoMessage() {}

func (x *KindRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_UserKind_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KindRegisterResponse.ProtoReflect.Descriptor instead.
func (*KindRegisterResponse) Descriptor() ([]byte, []int) {
	return file_UserKind_proto_rawDescGZIP(), []int{3}
}

func (x *KindRegisterResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *KindRegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *KindRegisterResponse) GetData() *UserModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_UserKind_proto protoreflect.FileDescriptor

var file_UserKind_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37,
	0x0a, 0x10, 0x4b, 0x69, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x61, 0x0a, 0x11, 0x4b, 0x69, 0x6e, 0x64, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x13, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x64, 0x0a, 0x14, 0x4b, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x15, 0x5a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UserKind_proto_rawDescOnce sync.Once
	file_UserKind_proto_rawDescData = file_UserKind_proto_rawDesc
)

func file_UserKind_proto_rawDescGZIP() []byte {
	file_UserKind_proto_rawDescOnce.Do(func() {
		file_UserKind_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserKind_proto_rawDescData)
	})
	return file_UserKind_proto_rawDescData
}

var file_UserKind_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_UserKind_proto_goTypes = []interface{}{
	(*KindLoginRequest)(nil),     // 0: KindLoginRequest
	(*KindLoginResponse)(nil),    // 1: KindLoginResponse
	(*KindRegisterRequest)(nil),  // 2: KindRegisterRequest
	(*KindRegisterResponse)(nil), // 3: KindRegisterResponse
	(*UserLoginModel)(nil),       // 4: UserLoginModel
	(*UserModel)(nil),            // 5: UserModel
	(*UserRegisterModel)(nil),    // 6: UserRegisterModel
}
var file_UserKind_proto_depIdxs = []int32{
	4, // 0: KindLoginRequest.spec:type_name -> UserLoginModel
	5, // 1: KindLoginResponse.data:type_name -> UserModel
	6, // 2: KindRegisterRequest.spec:type_name -> UserRegisterModel
	5, // 3: KindRegisterResponse.data:type_name -> UserModel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_UserKind_proto_init() }
func file_UserKind_proto_init() {
	if File_UserKind_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UserKind_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KindLoginRequest); i {
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
		file_UserKind_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KindLoginResponse); i {
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
		file_UserKind_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KindRegisterRequest); i {
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
		file_UserKind_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KindRegisterResponse); i {
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
			RawDescriptor: file_UserKind_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UserKind_proto_goTypes,
		DependencyIndexes: file_UserKind_proto_depIdxs,
		MessageInfos:      file_UserKind_proto_msgTypes,
	}.Build()
	File_UserKind_proto = out.File
	file_UserKind_proto_rawDesc = nil
	file_UserKind_proto_goTypes = nil
	file_UserKind_proto_depIdxs = nil
}
