// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: v2ray.com/core/proxy/vless/inbound/config.proto

package inbound

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	net "v2ray.com/core/common/net"
	protocol "v2ray.com/core/common/protocol"
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

type Fallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr *net.IPOrDomain `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Unix string          `protobuf:"bytes,3,opt,name=unix,proto3" json:"unix,omitempty"`
}

func (x *Fallback) Reset() {
	*x = Fallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fallback) ProtoMessage() {}

func (x *Fallback) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fallback.ProtoReflect.Descriptor instead.
func (*Fallback) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP(), []int{0}
}

func (x *Fallback) GetAddr() *net.IPOrDomain {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *Fallback) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Fallback) GetUnix() string {
	if x != nil {
		return x.Unix
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*protocol.User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	// Decryption settings. Only applies to server side, and only accepts "none" for now.
	Decryption string    `protobuf:"bytes,2,opt,name=decryption,proto3" json:"decryption,omitempty"`
	Fallback   *Fallback `protobuf:"bytes,3,opt,name=fallback,proto3" json:"fallback,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetUser() []*protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Config) GetDecryption() string {
	if x != nil {
		return x.Decryption
	}
	return ""
}

func (x *Config) GetFallback() *Fallback {
	if x != nil {
		return x.Fallback
	}
	return nil
}

var File_v2ray_com_core_proxy_vless_inbound_config_proto protoreflect.FileDescriptor

var file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x69, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x1a, 0x27, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x08, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x12, 0x35, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x78,
	0x22, 0xa4, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x50, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x50, 0x01, 0x5a,
	0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0xaa, 0x02, 0x1e, 0x56, 0x32, 0x52, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x56, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescOnce sync.Once
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData = file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc
)

func file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData)
	})
	return file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDescData
}

var file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes = []interface{}{
	(*Fallback)(nil),       // 0: v2ray.core.proxy.vless.inbound.Fallback
	(*Config)(nil),         // 1: v2ray.core.proxy.vless.inbound.Config
	(*net.IPOrDomain)(nil), // 2: v2ray.core.common.net.IPOrDomain
	(*protocol.User)(nil),  // 3: v2ray.core.common.protocol.User
}
var file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.proxy.vless.inbound.Fallback.addr:type_name -> v2ray.core.common.net.IPOrDomain
	3, // 1: v2ray.core.proxy.vless.inbound.Config.user:type_name -> v2ray.core.common.protocol.User
	0, // 2: v2ray.core.proxy.vless.inbound.Config.fallback:type_name -> v2ray.core.proxy.vless.inbound.Fallback
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_proxy_vless_inbound_config_proto_init() }
func file_v2ray_com_core_proxy_vless_inbound_config_proto_init() {
	if File_v2ray_com_core_proxy_vless_inbound_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fallback); i {
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
		file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_proxy_vless_inbound_config_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_proxy_vless_inbound_config_proto = out.File
	file_v2ray_com_core_proxy_vless_inbound_config_proto_rawDesc = nil
	file_v2ray_com_core_proxy_vless_inbound_config_proto_goTypes = nil
	file_v2ray_com_core_proxy_vless_inbound_config_proto_depIdxs = nil
}
