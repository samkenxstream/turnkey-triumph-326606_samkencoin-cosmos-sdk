// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/gov/v1beta1/genesis.proto

package govv1beta1

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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// deposits defines all the deposits present at genesis.
	Deposits []*Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits,omitempty"`
	// votes defines all the votes present at genesis.
	Votes []*Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
	// proposals defines all the proposals present at genesis.
	Proposals []*Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// params defines all the paramaters of related to deposit.
	DepositParams *DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params,omitempty"`
	// params defines all the paramaters of related to voting.
	VotingParams *VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params,omitempty"`
	// params defines all the paramaters of related to tally.
	TallyParams *TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetStartingProposalId() uint64 {
	if x != nil {
		return x.StartingProposalId
	}
	return 0
}

func (x *GenesisState) GetDeposits() []*Deposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

func (x *GenesisState) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *GenesisState) GetProposals() []*Proposal {
	if x != nil {
		return x.Proposals
	}
	return nil
}

func (x *GenesisState) GetDepositParams() *DepositParams {
	if x != nil {
		return x.DepositParams
	}
	return nil
}

func (x *GenesisState) GetVotingParams() *VotingParams {
	if x != nil {
		return x.VotingParams
	}
	return nil
}

func (x *GenesisState) GetTallyParams() *TallyParams {
	if x != nil {
		return x.TallyParams
	}
	return nil
}

var File_cosmos_gov_v1beta1_genesis_proto protoreflect.FileDescriptor

var file_cosmos_gov_v1beta1_genesis_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x67, 0x6f, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04, 0x0a, 0x0c, 0x47,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x49, 0x0a,
	0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x42, 0x10, 0xc8, 0xde,
	0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x08, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x52, 0x08,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x42, 0x0d, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x05, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x42, 0x11, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf,
	0x1f, 0x09, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x74, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54,
	0x61, 0x6c, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0b, 0x74, 0x61, 0x6c, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0xd0, 0x01,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x67, 0x6f,
	0x76, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x47, 0x58, 0xaa, 0x02,
	0x12, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x47, 0x6f, 0x76, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x47, 0x6f, 0x76,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x47, 0x6f, 0x76, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x76, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_gov_v1beta1_genesis_proto_rawDescOnce sync.Once
	file_cosmos_gov_v1beta1_genesis_proto_rawDescData = file_cosmos_gov_v1beta1_genesis_proto_rawDesc
)

func file_cosmos_gov_v1beta1_genesis_proto_rawDescGZIP() []byte {
	file_cosmos_gov_v1beta1_genesis_proto_rawDescOnce.Do(func() {
		file_cosmos_gov_v1beta1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_gov_v1beta1_genesis_proto_rawDescData)
	})
	return file_cosmos_gov_v1beta1_genesis_proto_rawDescData
}

var file_cosmos_gov_v1beta1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_gov_v1beta1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),  // 0: cosmos.gov.v1beta1.GenesisState
	(*Deposit)(nil),       // 1: cosmos.gov.v1beta1.Deposit
	(*Vote)(nil),          // 2: cosmos.gov.v1beta1.Vote
	(*Proposal)(nil),      // 3: cosmos.gov.v1beta1.Proposal
	(*DepositParams)(nil), // 4: cosmos.gov.v1beta1.DepositParams
	(*VotingParams)(nil),  // 5: cosmos.gov.v1beta1.VotingParams
	(*TallyParams)(nil),   // 6: cosmos.gov.v1beta1.TallyParams
}
var file_cosmos_gov_v1beta1_genesis_proto_depIdxs = []int32{
	1, // 0: cosmos.gov.v1beta1.GenesisState.deposits:type_name -> cosmos.gov.v1beta1.Deposit
	2, // 1: cosmos.gov.v1beta1.GenesisState.votes:type_name -> cosmos.gov.v1beta1.Vote
	3, // 2: cosmos.gov.v1beta1.GenesisState.proposals:type_name -> cosmos.gov.v1beta1.Proposal
	4, // 3: cosmos.gov.v1beta1.GenesisState.deposit_params:type_name -> cosmos.gov.v1beta1.DepositParams
	5, // 4: cosmos.gov.v1beta1.GenesisState.voting_params:type_name -> cosmos.gov.v1beta1.VotingParams
	6, // 5: cosmos.gov.v1beta1.GenesisState.tally_params:type_name -> cosmos.gov.v1beta1.TallyParams
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cosmos_gov_v1beta1_genesis_proto_init() }
func file_cosmos_gov_v1beta1_genesis_proto_init() {
	if File_cosmos_gov_v1beta1_genesis_proto != nil {
		return
	}
	file_cosmos_gov_v1beta1_gov_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_gov_v1beta1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_cosmos_gov_v1beta1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_gov_v1beta1_genesis_proto_goTypes,
		DependencyIndexes: file_cosmos_gov_v1beta1_genesis_proto_depIdxs,
		MessageInfos:      file_cosmos_gov_v1beta1_genesis_proto_msgTypes,
	}.Build()
	File_cosmos_gov_v1beta1_genesis_proto = out.File
	file_cosmos_gov_v1beta1_genesis_proto_rawDesc = nil
	file_cosmos_gov_v1beta1_genesis_proto_goTypes = nil
	file_cosmos_gov_v1beta1_genesis_proto_depIdxs = nil
}
