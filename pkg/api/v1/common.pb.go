// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/proto/v1/common.proto

package v1

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

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ExternalAuthn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host *Host `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *ExternalAuthn) Reset() {
	*x = ExternalAuthn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAuthn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAuthn) ProtoMessage() {}

func (x *ExternalAuthn) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAuthn.ProtoReflect.Descriptor instead.
func (*ExternalAuthn) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalAuthn) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

type DirectAuthn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     *Host  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DirectAuthn) Reset() {
	*x = DirectAuthn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectAuthn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectAuthn) ProtoMessage() {}

func (x *DirectAuthn) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectAuthn.ProtoReflect.Descriptor instead.
func (*DirectAuthn) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *DirectAuthn) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *DirectAuthn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DirectAuthn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Authn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Authn:
	//	*Authn_DirectAuthn
	//	*Authn_ExternalAuthn
	Authn isAuthn_Authn `protobuf_oneof:"authn"`
}

func (x *Authn) Reset() {
	*x = Authn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authn) ProtoMessage() {}

func (x *Authn) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authn.ProtoReflect.Descriptor instead.
func (*Authn) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_common_proto_rawDescGZIP(), []int{3}
}

func (m *Authn) GetAuthn() isAuthn_Authn {
	if m != nil {
		return m.Authn
	}
	return nil
}

func (x *Authn) GetDirectAuthn() *DirectAuthn {
	if x, ok := x.GetAuthn().(*Authn_DirectAuthn); ok {
		return x.DirectAuthn
	}
	return nil
}

func (x *Authn) GetExternalAuthn() *ExternalAuthn {
	if x, ok := x.GetAuthn().(*Authn_ExternalAuthn); ok {
		return x.ExternalAuthn
	}
	return nil
}

type isAuthn_Authn interface {
	isAuthn_Authn()
}

type Authn_DirectAuthn struct {
	DirectAuthn *DirectAuthn `protobuf:"bytes,1,opt,name=directAuthn,proto3,oneof"`
}

type Authn_ExternalAuthn struct {
	ExternalAuthn *ExternalAuthn `protobuf:"bytes,2,opt,name=externalAuthn,proto3,oneof"`
}

func (*Authn_DirectAuthn) isAuthn_Authn() {}

func (*Authn_ExternalAuthn) isAuthn_Authn() {}

type Vendor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Vendor) Reset() {
	*x = Vendor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vendor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vendor) ProtoMessage() {}

func (x *Vendor) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vendor.ProtoReflect.Descriptor instead.
func (*Vendor) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *Vendor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_proto_v1_common_proto protoreflect.FileDescriptor

var file_api_proto_v1_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x22, 0x52, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68,
	0x6e, 0x12, 0x41, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x12, 0x41, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0xca, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x58, 0x0a, 0x0b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x12, 0x5e, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74,
	0x68, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x22, 0x1c, 0x0a, 0x06,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62,
	0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x6e, 0x6a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_common_proto_rawDescOnce sync.Once
	file_api_proto_v1_common_proto_rawDescData = file_api_proto_v1_common_proto_rawDesc
)

func file_api_proto_v1_common_proto_rawDescGZIP() []byte {
	file_api_proto_v1_common_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_common_proto_rawDescData)
	})
	return file_api_proto_v1_common_proto_rawDescData
}

var file_api_proto_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_v1_common_proto_goTypes = []interface{}{
	(*Host)(nil),          // 0: github.com.tinkerbell.pbnj.api.proto.v1.Host
	(*ExternalAuthn)(nil), // 1: github.com.tinkerbell.pbnj.api.proto.v1.ExternalAuthn
	(*DirectAuthn)(nil),   // 2: github.com.tinkerbell.pbnj.api.proto.v1.DirectAuthn
	(*Authn)(nil),         // 3: github.com.tinkerbell.pbnj.api.proto.v1.Authn
	(*Vendor)(nil),        // 4: github.com.tinkerbell.pbnj.api.proto.v1.Vendor
}
var file_api_proto_v1_common_proto_depIdxs = []int32{
	0, // 0: github.com.tinkerbell.pbnj.api.proto.v1.ExternalAuthn.host:type_name -> github.com.tinkerbell.pbnj.api.proto.v1.Host
	0, // 1: github.com.tinkerbell.pbnj.api.proto.v1.DirectAuthn.host:type_name -> github.com.tinkerbell.pbnj.api.proto.v1.Host
	2, // 2: github.com.tinkerbell.pbnj.api.proto.v1.Authn.directAuthn:type_name -> github.com.tinkerbell.pbnj.api.proto.v1.DirectAuthn
	1, // 3: github.com.tinkerbell.pbnj.api.proto.v1.Authn.externalAuthn:type_name -> github.com.tinkerbell.pbnj.api.proto.v1.ExternalAuthn
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_v1_common_proto_init() }
func file_api_proto_v1_common_proto_init() {
	if File_api_proto_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_api_proto_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAuthn); i {
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
		file_api_proto_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectAuthn); i {
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
		file_api_proto_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authn); i {
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
		file_api_proto_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vendor); i {
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
	file_api_proto_v1_common_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Authn_DirectAuthn)(nil),
		(*Authn_ExternalAuthn)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1_common_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_common_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_common_proto_msgTypes,
	}.Build()
	File_api_proto_v1_common_proto = out.File
	file_api_proto_v1_common_proto_rawDesc = nil
	file_api_proto_v1_common_proto_goTypes = nil
	file_api_proto_v1_common_proto_depIdxs = nil
}
