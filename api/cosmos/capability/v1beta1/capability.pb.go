// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/capability/v1beta1/capability.proto

package capabilityv1beta1

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

// Capability defines an implementation of an object capability. The index
// provided to a Capability must be globally unique.
type Capability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Capability) Reset() {
	*x = Capability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capability) ProtoMessage() {}

func (x *Capability) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capability.ProtoReflect.Descriptor instead.
func (*Capability) Descriptor() ([]byte, []int) {
	return file_cosmos_capability_v1beta1_capability_proto_rawDescGZIP(), []int{0}
}

func (x *Capability) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

// Owner defines a single capability owner. An owner is defined by the name of
// capability and the module name.
type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_cosmos_capability_v1beta1_capability_proto_rawDescGZIP(), []int{1}
}

func (x *Owner) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Owner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// CapabilityOwners defines a set of owners of a single Capability. The set of
// owners must be unique.
type CapabilityOwners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owners []*Owner `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
}

func (x *CapabilityOwners) Reset() {
	*x = CapabilityOwners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityOwners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityOwners) ProtoMessage() {}

func (x *CapabilityOwners) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_capability_v1beta1_capability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityOwners.ProtoReflect.Descriptor instead.
func (*CapabilityOwners) Descriptor() ([]byte, []int) {
	return file_cosmos_capability_v1beta1_capability_proto_rawDescGZIP(), []int{2}
}

func (x *CapabilityOwners) GetOwners() []*Owner {
	if x != nil {
		return x.Owners
	}
	return nil
}

var File_cosmos_capability_v1beta1_capability_proto protoreflect.FileDescriptor

var file_cosmos_capability_v1beta1_capability_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0x3d, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x08, 0x88, 0xa0,
	0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0x52, 0x0a, 0x10, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x42, 0x84, 0x02, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x43, 0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_capability_v1beta1_capability_proto_rawDescOnce sync.Once
	file_cosmos_capability_v1beta1_capability_proto_rawDescData = file_cosmos_capability_v1beta1_capability_proto_rawDesc
)

func file_cosmos_capability_v1beta1_capability_proto_rawDescGZIP() []byte {
	file_cosmos_capability_v1beta1_capability_proto_rawDescOnce.Do(func() {
		file_cosmos_capability_v1beta1_capability_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_capability_v1beta1_capability_proto_rawDescData)
	})
	return file_cosmos_capability_v1beta1_capability_proto_rawDescData
}

var file_cosmos_capability_v1beta1_capability_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_capability_v1beta1_capability_proto_goTypes = []interface{}{
	(*Capability)(nil),       // 0: cosmos.capability.v1beta1.Capability
	(*Owner)(nil),            // 1: cosmos.capability.v1beta1.Owner
	(*CapabilityOwners)(nil), // 2: cosmos.capability.v1beta1.CapabilityOwners
}
var file_cosmos_capability_v1beta1_capability_proto_depIdxs = []int32{
	1, // 0: cosmos.capability.v1beta1.CapabilityOwners.owners:type_name -> cosmos.capability.v1beta1.Owner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_capability_v1beta1_capability_proto_init() }
func file_cosmos_capability_v1beta1_capability_proto_init() {
	if File_cosmos_capability_v1beta1_capability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_capability_v1beta1_capability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capability); i {
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
		file_cosmos_capability_v1beta1_capability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
		file_cosmos_capability_v1beta1_capability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityOwners); i {
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
			RawDescriptor: file_cosmos_capability_v1beta1_capability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_capability_v1beta1_capability_proto_goTypes,
		DependencyIndexes: file_cosmos_capability_v1beta1_capability_proto_depIdxs,
		MessageInfos:      file_cosmos_capability_v1beta1_capability_proto_msgTypes,
	}.Build()
	File_cosmos_capability_v1beta1_capability_proto = out.File
	file_cosmos_capability_v1beta1_capability_proto_rawDesc = nil
	file_cosmos_capability_v1beta1_capability_proto_goTypes = nil
	file_cosmos_capability_v1beta1_capability_proto_depIdxs = nil
}
