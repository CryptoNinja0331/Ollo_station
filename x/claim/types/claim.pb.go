// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/claim/v1/claim.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClaimAction int32

const (
	//
	ClaimActionUnspecified ClaimAction = 0
	//
	ClaimActionInitialClaim ClaimAction = 1
	//
	ClaimActionStake ClaimAction = 2
	//
	ClaimActionCastVote ClaimAction = 3
	//
	ClaimActionDepositLiquidity ClaimAction = 4
	//
	ClaimActionLockLiquidity ClaimAction = 5
	//
	ClaimActionMintNft ClaimAction = 6
	//
	ClaimActionAuctionNft ClaimAction = 7
	//
	ClaimActionDepositExternalAsset ClaimAction = 8
	//
	ClaimActionSwap ClaimAction = 9
	//
	ClaimActionBorrow ClaimAction = 10
)

var ClaimAction_name = map[int32]string{
	0:  "ClaimActionUnspecified",
	1:  "ClaimActionInitialClaim",
	2:  "ClaimActionStake",
	3:  "ClaimActionCastVote",
	4:  "ClaimActionDepositLiquidity",
	5:  "ClaimActionLockLiquidity",
	6:  "ClaimActionMintNft",
	7:  "ClaimActionAuctionNft",
	8:  "ClaimActionDepositExternalAsset",
	9:  "ClaimActionSwap",
	10: "ClaimActionBorrow",
}

var ClaimAction_value = map[string]int32{
	"ClaimActionUnspecified":          0,
	"ClaimActionInitialClaim":         1,
	"ClaimActionStake":                2,
	"ClaimActionCastVote":             3,
	"ClaimActionDepositLiquidity":     4,
	"ClaimActionLockLiquidity":        5,
	"ClaimActionMintNft":              6,
	"ClaimActionAuctionNft":           7,
	"ClaimActionDepositExternalAsset": 8,
	"ClaimActionSwap":                 9,
	"ClaimActionBorrow":               10,
}

func (x ClaimAction) String() string {
	return proto.EnumName(ClaimAction_name, int32(x))
}

func (ClaimAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{0}
}

type ClaimStatus int32

const (
	//
	ClaimStatusUnspecified ClaimStatus = 0
	//
	ClaimStatusPending ClaimStatus = 1
	//
	ClaimStatusClaimed ClaimStatus = 2
	//
	ClaimStatusActive ClaimStatus = 3
	//
	ClaimStatusFailed ClaimStatus = 4
)

var ClaimStatus_name = map[int32]string{
	0: "ClaimStatusUnspecified",
	1: "ClaimStatusPending",
	2: "ClaimStatusClaimed",
	3: "ClaimStatusActive",
	4: "ClaimStatusExpired",
}

var ClaimStatus_value = map[string]int32{
	"ClaimStatusUnspecified": 0,
	"ClaimStatusPending":     1,
	"ClaimStatusClaimed":     2,
	"ClaimStatusActive":      3,
	"ClaimStatusExpired":     4,
}

func (x ClaimStatus) String() string {
	return proto.EnumName(ClaimStatus_name, int32(x))
}

func (ClaimStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{1}
}

type InitialClaim struct {
	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" yaml:"enabled"`
	GoalId  uint64 `protobuf:"varint,2,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty" yaml:"goal_id"`
}

func (m *InitialClaim) Reset()         { *m = InitialClaim{} }
func (m *InitialClaim) String() string { return proto.CompactTextString(m) }
func (*InitialClaim) ProtoMessage()    {}
func (*InitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{0}
}
func (m *InitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialClaim.Merge(m, src)
}
func (m *InitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *InitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_InitialClaim proto.InternalMessageInfo

type ClaimRecord struct {
	Address        string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Claimable      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=claimable,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimable" yaml:"claimable"`
	CompletedGoals []uint64                               `protobuf:"varint,3,rep,packed,name=completed_goals,json=completedGoals,proto3" json:"completed_goals,omitempty" yaml:"completed_goals"`
	ClaimedGoals   []uint64                               `protobuf:"varint,4,rep,packed,name=claimed_goals,json=claimedGoals,proto3" json:"claimed_goals,omitempty" yaml:"claimed_goals"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcf5997285c872eb, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.claim.v1.ClaimAction", ClaimAction_name, ClaimAction_value)
	proto.RegisterEnum("ollo.claim.v1.ClaimStatus", ClaimStatus_name, ClaimStatus_value)
	proto.RegisterType((*InitialClaim)(nil), "ollo.claim.v1.InitialClaim")
	proto.RegisterType((*ClaimRecord)(nil), "ollo.claim.v1.ClaimRecord")
}

func init() { proto.RegisterFile("ollo/claim/v1/claim.proto", fileDescriptor_fcf5997285c872eb) }

var fileDescriptor_fcf5997285c872eb = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0xed, 0x90, 0x4b, 0xc8, 0x5c, 0x3e, 0x7c, 0x87, 0x00, 0x61, 0xb8, 0xb2, 0xad, 0x5c,
	0xe9, 0x2a, 0x4a, 0x49, 0x52, 0xd4, 0xaa, 0xaa, 0x90, 0x2a, 0x35, 0x09, 0xb4, 0x8d, 0x44, 0x2b,
	0x14, 0xda, 0x2e, 0xba, 0x41, 0xc6, 0x1e, 0xd2, 0x11, 0x8e, 0x27, 0xf5, 0x4c, 0xf8, 0x78, 0x03,
	0xe4, 0x55, 0x5f, 0xc0, 0xab, 0x6e, 0xfa, 0x00, 0x7d, 0x87, 0xb2, 0x44, 0x5d, 0x55, 0x5d, 0x44,
	0x10, 0x56, 0xdd, 0xe6, 0x09, 0x2a, 0x7f, 0x90, 0x0c, 0x84, 0xb4, 0x2b, 0xcf, 0xfc, 0x7f, 0xe7,
	0x63, 0xce, 0xd1, 0xf1, 0x01, 0xcb, 0xd4, 0xb6, 0x69, 0xd9, 0xb4, 0x0d, 0xd2, 0x2a, 0x1f, 0xae,
	0x45, 0x87, 0x52, 0xdb, 0xa5, 0x9c, 0xc2, 0x99, 0x00, 0x95, 0x22, 0xe5, 0x70, 0x0d, 0x65, 0x9a,
	0xb4, 0x49, 0x43, 0x52, 0x0e, 0x4e, 0x91, 0x11, 0x5a, 0x36, 0x29, 0x6b, 0x51, 0xb6, 0x1b, 0x81,
	0xe8, 0x12, 0xa1, 0x1c, 0x01, 0xd3, 0x75, 0x87, 0x70, 0x62, 0xd8, 0xb5, 0x20, 0x06, 0x5c, 0x05,
	0x29, 0xec, 0x18, 0x7b, 0x36, 0xb6, 0xb2, 0xb2, 0x2e, 0xe7, 0xa7, 0xaa, 0xb0, 0xdf, 0xd5, 0x66,
	0x4f, 0x8c, 0x96, 0xbd, 0x9e, 0x8b, 0x41, 0xae, 0x71, 0x6d, 0x02, 0xef, 0x81, 0x54, 0x93, 0x1a,
	0xf6, 0x2e, 0xb1, 0xb2, 0x09, 0x5d, 0xce, 0x27, 0x45, 0xeb, 0x18, 0xe4, 0x1a, 0x93, 0xc1, 0xa9,
	0x6e, 0xe5, 0xbe, 0x26, 0xc0, 0xdf, 0x61, 0x92, 0x06, 0x36, 0xa9, 0x6b, 0xc1, 0x0d, 0x90, 0x32,
	0x2c, 0xcb, 0xc5, 0x8c, 0x85, 0xa9, 0xd2, 0xd5, 0xc2, 0xd0, 0x39, 0x06, 0xb9, 0x6f, 0x5f, 0x8a,
	0x99, 0xf8, 0xbd, 0x95, 0x48, 0xda, 0xe1, 0x2e, 0x71, 0x9a, 0x8d, 0x6b, 0x57, 0xe8, 0x80, 0x74,
	0x58, 0x7d, 0xf0, 0xa0, 0xf0, 0x11, 0xe9, 0xea, 0xf6, 0x59, 0x57, 0x93, 0x7e, 0x74, 0xb5, 0xff,
	0x9b, 0x84, 0xbf, 0xef, 0xec, 0x95, 0x4c, 0xda, 0x8a, 0x8b, 0x8e, 0x3f, 0x45, 0x66, 0x1d, 0x94,
	0xf9, 0x49, 0x1b, 0xb3, 0x52, 0xdd, 0xe1, 0xfd, 0xae, 0xa6, 0x44, 0x59, 0x07, 0x81, 0x82, 0xbc,
	0x20, 0xce, 0x5b, 0x77, 0x78, 0x63, 0x98, 0x02, 0xd6, 0xc0, 0x9c, 0x49, 0x5b, 0x6d, 0x1b, 0x73,
	0x6c, 0xed, 0x06, 0x95, 0xb1, 0xec, 0x84, 0x3e, 0x91, 0x4f, 0x56, 0x51, 0xbf, 0xab, 0x2d, 0xc6,
	0x71, 0x6e, 0x1a, 0xe4, 0x1a, 0xb3, 0x03, 0xe5, 0x79, 0x20, 0xc0, 0x27, 0x60, 0x26, 0x8c, 0x38,
	0x08, 0x91, 0x0c, 0x43, 0x64, 0xfb, 0x5d, 0x2d, 0x23, 0x3c, 0x65, 0x18, 0x60, 0x3a, 0xbe, 0x87,
	0xee, 0x85, 0x9f, 0xc9, 0xb8, 0x93, 0x15, 0x93, 0x13, 0xea, 0xc0, 0x47, 0x60, 0x51, 0xb8, 0xbe,
	0x71, 0x58, 0x1b, 0x9b, 0x64, 0x9f, 0x60, 0x4b, 0x91, 0x10, 0xf2, 0x7c, 0x7d, 0x0c, 0x85, 0x8f,
	0xc1, 0x92, 0x40, 0xc4, 0x39, 0x50, 0x64, 0xb4, 0xe2, 0xf9, 0xfa, 0x38, 0x0c, 0x0b, 0x40, 0x11,
	0xd0, 0x0e, 0x37, 0x0e, 0xb0, 0x92, 0x40, 0x19, 0xcf, 0xd7, 0x47, 0x74, 0x78, 0x1f, 0xcc, 0x0b,
	0x5a, 0xcd, 0x60, 0xfc, 0x2d, 0xe5, 0x58, 0x99, 0x40, 0x4b, 0x9e, 0xaf, 0xdf, 0x85, 0xe0, 0x53,
	0xb0, 0x22, 0xc8, 0x1b, 0xb8, 0x4d, 0x19, 0xe1, 0x5b, 0xe4, 0x43, 0x87, 0x58, 0x84, 0x9f, 0x28,
	0x49, 0xa4, 0x79, 0xbe, 0xfe, 0x3b, 0x13, 0xb8, 0x0e, 0xb2, 0x02, 0xde, 0xa2, 0xe6, 0xc1, 0xd0,
	0xfd, 0x2f, 0xf4, 0xaf, 0xe7, 0xeb, 0x63, 0x39, 0x2c, 0x01, 0x28, 0xb0, 0x97, 0xc4, 0xe1, 0xaf,
	0xf6, 0xb9, 0x32, 0x89, 0x16, 0x3d, 0x5f, 0xbf, 0x83, 0xc0, 0x87, 0x60, 0x41, 0x50, 0x2b, 0x9d,
	0xf0, 0x13, 0xb8, 0xa4, 0xd0, 0xb2, 0xe7, 0xeb, 0x77, 0x43, 0xf8, 0x02, 0x68, 0xa3, 0x05, 0x6c,
	0x1e, 0x73, 0xec, 0x3a, 0x86, 0x5d, 0x61, 0x0c, 0x73, 0x65, 0x0a, 0xfd, 0xe7, 0xf9, 0xfa, 0x9f,
	0xcc, 0x60, 0x1e, 0xcc, 0x89, 0x3d, 0x3f, 0x32, 0xda, 0x4a, 0x1a, 0xcd, 0x7b, 0xbe, 0x7e, 0x5b,
	0x86, 0xab, 0xe0, 0x1f, 0x41, 0xaa, 0x52, 0xd7, 0xa5, 0x47, 0x0a, 0x40, 0x0b, 0x9e, 0xaf, 0x8f,
	0x02, 0x94, 0x3c, 0xfd, 0xa4, 0x4a, 0x85, 0xd3, 0xeb, 0xbf, 0x76, 0x87, 0x1b, 0xbc, 0xc3, 0x06,
	0xb3, 0x16, 0x5d, 0xc7, 0xcd, 0xda, 0x08, 0x1d, 0x74, 0x35, 0x22, 0xdb, 0xd8, 0xb1, 0x88, 0xd3,
	0x54, 0x64, 0xa1, 0xab, 0x37, 0xc8, 0x2d, 0xfb, 0x5a, 0x34, 0xfe, 0x4a, 0x62, 0xc4, 0x3e, 0x26,
	0x83, 0xda, 0x22, 0x35, 0x28, 0xe4, 0x30, 0x98, 0xb1, 0x61, 0x6d, 0x22, 0x80, 0xc5, 0x1b, 0xd1,
	0x37, 0x8f, 0xdb, 0xc4, 0xc5, 0x96, 0x92, 0x1c, 0x31, 0x7f, 0x66, 0x10, 0x1b, 0x5b, 0x51, 0x2b,
	0xaa, 0xaf, 0xcf, 0x2e, 0x55, 0xe9, 0xfc, 0x52, 0x95, 0x2e, 0x2e, 0x55, 0xf9, 0x73, 0x4f, 0x95,
	0xcf, 0x7a, 0xaa, 0x7c, 0xde, 0x53, 0xe5, 0x8b, 0x9e, 0x2a, 0x7f, 0xbc, 0x52, 0xa5, 0xf3, 0x2b,
	0x55, 0xfa, 0x7e, 0xa5, 0x4a, 0xef, 0x0a, 0xc2, 0xe6, 0x09, 0x16, 0x74, 0x91, 0x71, 0x23, 0xe8,
	0x6c, 0x78, 0x29, 0x1f, 0xc7, 0xab, 0x3c, 0xdc, 0x40, 0x7b, 0x93, 0xe1, 0x22, 0x7e, 0xf0, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0xed, 0x11, 0x26, 0xe5, 0x05, 0x00, 0x00,
}

func (this *InitialClaim) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*InitialClaim)
	if !ok {
		that2, ok := that.(InitialClaim)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *InitialClaim")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *InitialClaim but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *InitialClaim but is not nil && this == nil")
	}
	if this.Enabled != that1.Enabled {
		return fmt.Errorf("Enabled this(%v) Not Equal that(%v)", this.Enabled, that1.Enabled)
	}
	if this.GoalId != that1.GoalId {
		return fmt.Errorf("GoalId this(%v) Not Equal that(%v)", this.GoalId, that1.GoalId)
	}
	return nil
}
func (this *InitialClaim) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InitialClaim)
	if !ok {
		that2, ok := that.(InitialClaim)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.GoalId != that1.GoalId {
		return false
	}
	return true
}
func (this *ClaimRecord) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClaimRecord)
	if !ok {
		that2, ok := that.(ClaimRecord)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ClaimRecord")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClaimRecord but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClaimRecord but is not nil && this == nil")
	}
	if this.Address != that1.Address {
		return fmt.Errorf("Address this(%v) Not Equal that(%v)", this.Address, that1.Address)
	}
	if !this.Claimable.Equal(that1.Claimable) {
		return fmt.Errorf("Claimable this(%v) Not Equal that(%v)", this.Claimable, that1.Claimable)
	}
	if len(this.CompletedGoals) != len(that1.CompletedGoals) {
		return fmt.Errorf("CompletedGoals this(%v) Not Equal that(%v)", len(this.CompletedGoals), len(that1.CompletedGoals))
	}
	for i := range this.CompletedGoals {
		if this.CompletedGoals[i] != that1.CompletedGoals[i] {
			return fmt.Errorf("CompletedGoals this[%v](%v) Not Equal that[%v](%v)", i, this.CompletedGoals[i], i, that1.CompletedGoals[i])
		}
	}
	if len(this.ClaimedGoals) != len(that1.ClaimedGoals) {
		return fmt.Errorf("ClaimedGoals this(%v) Not Equal that(%v)", len(this.ClaimedGoals), len(that1.ClaimedGoals))
	}
	for i := range this.ClaimedGoals {
		if this.ClaimedGoals[i] != that1.ClaimedGoals[i] {
			return fmt.Errorf("ClaimedGoals this[%v](%v) Not Equal that[%v](%v)", i, this.ClaimedGoals[i], i, that1.ClaimedGoals[i])
		}
	}
	return nil
}
func (this *ClaimRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimRecord)
	if !ok {
		that2, ok := that.(ClaimRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !this.Claimable.Equal(that1.Claimable) {
		return false
	}
	if len(this.CompletedGoals) != len(that1.CompletedGoals) {
		return false
	}
	for i := range this.CompletedGoals {
		if this.CompletedGoals[i] != that1.CompletedGoals[i] {
			return false
		}
	}
	if len(this.ClaimedGoals) != len(that1.ClaimedGoals) {
		return false
	}
	for i := range this.ClaimedGoals {
		if this.ClaimedGoals[i] != that1.ClaimedGoals[i] {
			return false
		}
	}
	return true
}
func (m *InitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GoalId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.GoalId))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedGoals) > 0 {
		dAtA2 := make([]byte, len(m.ClaimedGoals)*10)
		var j1 int
		for _, num := range m.ClaimedGoals {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintClaim(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompletedGoals) > 0 {
		dAtA4 := make([]byte, len(m.CompletedGoals)*10)
		var j3 int
		for _, num := range m.CompletedGoals {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaim(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Claimable.Size()
		i -= size
		if _, err := m.Claimable.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.GoalId != 0 {
		n += 1 + sovClaim(uint64(m.GoalId))
	}
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.Claimable.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.CompletedGoals) > 0 {
		l = 0
		for _, e := range m.CompletedGoals {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	if len(m.ClaimedGoals) > 0 {
		l = 0
		for _, e := range m.ClaimedGoals {
			l += sovClaim(uint64(e))
		}
		n += 1 + sovClaim(uint64(l)) + l
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InitialClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalId", wireType)
			}
			m.GoalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GoalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Claimable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CompletedGoals = append(m.CompletedGoals, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CompletedGoals) == 0 {
					m.CompletedGoals = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CompletedGoals = append(m.CompletedGoals, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedGoals", wireType)
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedGoals = append(m.ClaimedGoals, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ClaimedGoals) == 0 {
					m.ClaimedGoals = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedGoals = append(m.ClaimedGoals, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedGoals", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)