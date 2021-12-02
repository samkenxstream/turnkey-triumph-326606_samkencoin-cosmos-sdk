// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/base/store/v1beta1/commit_info.proto

package storev1beta1

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

// CommitInfo defines commit information used by the multi-store when committing
// a version/height.
type CommitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    int64        `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	StoreInfos []*StoreInfo `protobuf:"bytes,2,rep,name=store_infos,json=storeInfos,proto3" json:"store_infos,omitempty"`
}

func (x *CommitInfo) Reset() {
	*x = CommitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitInfo) ProtoMessage() {}

func (x *CommitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitInfo.ProtoReflect.Descriptor instead.
func (*CommitInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_commit_info_proto_rawDescGZIP(), []int{0}
}

func (x *CommitInfo) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CommitInfo) GetStoreInfos() []*StoreInfo {
	if x != nil {
		return x.StoreInfos
	}
	return nil
}

// StoreInfo defines store-specific commit information. It contains a reference
// between a store name and the commit ID.
type StoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CommitId *CommitID `protobuf:"bytes,2,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
}

func (x *StoreInfo) Reset() {
	*x = StoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreInfo) ProtoMessage() {}

func (x *StoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreInfo.ProtoReflect.Descriptor instead.
func (*StoreInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_commit_info_proto_rawDescGZIP(), []int{1}
}

func (x *StoreInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreInfo) GetCommitId() *CommitID {
	if x != nil {
		return x.CommitId
	}
	return nil
}

// CommitID defines the committment information when a specific store is
// committed.
type CommitID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Hash    []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CommitID) Reset() {
	*x = CommitID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitID) ProtoMessage() {}

func (x *CommitID) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitID.ProtoReflect.Descriptor instead.
func (*CommitID) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_commit_info_proto_rawDescGZIP(), []int{2}
}

func (x *CommitID) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CommitID) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

var File_cosmos_base_store_v1beta1_commit_info_proto protoreflect.FileDescriptor

var file_cosmos_base_store_v1beta1_commit_info_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73,
	0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x42, 0x80, 0x02, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x53,
	0xaa, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x19, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x65, 0x3a,
	0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_store_v1beta1_commit_info_proto_rawDescOnce sync.Once
	file_cosmos_base_store_v1beta1_commit_info_proto_rawDescData = file_cosmos_base_store_v1beta1_commit_info_proto_rawDesc
)

func file_cosmos_base_store_v1beta1_commit_info_proto_rawDescGZIP() []byte {
	file_cosmos_base_store_v1beta1_commit_info_proto_rawDescOnce.Do(func() {
		file_cosmos_base_store_v1beta1_commit_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_store_v1beta1_commit_info_proto_rawDescData)
	})
	return file_cosmos_base_store_v1beta1_commit_info_proto_rawDescData
}

var file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_base_store_v1beta1_commit_info_proto_goTypes = []interface{}{
	(*CommitInfo)(nil), // 0: cosmos.base.store.v1beta1.CommitInfo
	(*StoreInfo)(nil),  // 1: cosmos.base.store.v1beta1.StoreInfo
	(*CommitID)(nil),   // 2: cosmos.base.store.v1beta1.CommitID
}
var file_cosmos_base_store_v1beta1_commit_info_proto_depIdxs = []int32{
	1, // 0: cosmos.base.store.v1beta1.CommitInfo.store_infos:type_name -> cosmos.base.store.v1beta1.StoreInfo
	2, // 1: cosmos.base.store.v1beta1.StoreInfo.commit_id:type_name -> cosmos.base.store.v1beta1.CommitID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_base_store_v1beta1_commit_info_proto_init() }
func file_cosmos_base_store_v1beta1_commit_info_proto_init() {
	if File_cosmos_base_store_v1beta1_commit_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitInfo); i {
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
		file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreInfo); i {
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
		file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitID); i {
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
			RawDescriptor: file_cosmos_base_store_v1beta1_commit_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_base_store_v1beta1_commit_info_proto_goTypes,
		DependencyIndexes: file_cosmos_base_store_v1beta1_commit_info_proto_depIdxs,
		MessageInfos:      file_cosmos_base_store_v1beta1_commit_info_proto_msgTypes,
	}.Build()
	File_cosmos_base_store_v1beta1_commit_info_proto = out.File
	file_cosmos_base_store_v1beta1_commit_info_proto_rawDesc = nil
	file_cosmos_base_store_v1beta1_commit_info_proto_goTypes = nil
	file_cosmos_base_store_v1beta1_commit_info_proto_depIdxs = nil
}
