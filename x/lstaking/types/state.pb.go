// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/lstaking/v1/state.proto

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

// data about supply and circulation of liquid token for calculation
type NetAmountState struct {
	// mint rate of liquid token, or total supply / net amount
	MintRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=mint_rate,json=mintRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mint_rate" yaml:"mint_rate"`
	// total supply of liquid token
	TotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_supply" yaml:"total_supply"`
	// net amount is sum of reserve, unbonding, rewards, and tokens
	NetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=net_amount,json=netAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"net_amount" yaml:"net_amount"`
	// total delegation shares of all validators
	TotalDelegationShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=total_delegation_shares,json=totalDelegationShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_delegation_shares" yaml:"total_delegation_shares"`
	// total worth of all delegation shares of all validators
	TotalLiquidTokens github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_liquid_tokens,json=totalLiquidTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_liquid_tokens" yaml:"total_liquid_tokens"`
	// total remaining rewards balance by all validators
	RewardsBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=rewards_balance,json=rewardsBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rewards_balance" yaml:"total_rewards_balance"`
	// total unbonding balance by all validators
	UnbondingBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=unbonding_balance,json=unbondingBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"unbonding_balance" yaml:"total_unbonding_balance"`
	// total liquid tokens in reserve account
	ReserveAccountBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=reserve_account_balance,json=reserveAccountBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserve_account_balance" yaml:"reserve_account_balance"`
}

func (m *NetAmountState) Reset()         { *m = NetAmountState{} }
func (m *NetAmountState) String() string { return proto.CompactTextString(m) }
func (*NetAmountState) ProtoMessage()    {}
func (*NetAmountState) Descriptor() ([]byte, []int) {
	return fileDescriptor_648524c7f61826f5, []int{0}
}
func (m *NetAmountState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetAmountState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetAmountState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetAmountState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAmountState.Merge(m, src)
}
func (m *NetAmountState) XXX_Size() int {
	return m.Size()
}
func (m *NetAmountState) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAmountState.DiscardUnknown(m)
}

var xxx_messageInfo_NetAmountState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NetAmountState)(nil), "ollo.lstaking.v1.NetAmountState")
}

func init() { proto.RegisterFile("ollo/lstaking/v1/state.proto", fileDescriptor_648524c7f61826f5) }

var fileDescriptor_648524c7f61826f5 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x31, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0xed, 0xff, 0x1f, 0xda, 0xe4, 0x40, 0x25, 0x71, 0xa9, 0x30, 0x51, 0x65, 0xa3, 0x1b,
	0x10, 0x4b, 0x6c, 0x55, 0x6c, 0xdd, 0x1a, 0x8a, 0x10, 0x52, 0x55, 0x21, 0x87, 0x89, 0xc5, 0x3a,
	0xdb, 0x27, 0xc7, 0xca, 0xf9, 0x2e, 0xf8, 0xce, 0x29, 0x91, 0x60, 0xef, 0xc8, 0x47, 0xe8, 0xc7,
	0xe9, 0xd8, 0x05, 0x09, 0x31, 0x44, 0x28, 0x59, 0x98, 0xfb, 0x09, 0xd0, 0xdd, 0x39, 0xae, 0x4b,
	0xc5, 0x60, 0x31, 0xe5, 0xde, 0xf7, 0x89, 0x9f, 0xdf, 0x73, 0xaf, 0x4e, 0x2f, 0xd8, 0x67, 0x84,
	0x30, 0x9f, 0x70, 0x81, 0xa6, 0x19, 0x4d, 0xfd, 0xf9, 0x81, 0xcf, 0x05, 0x12, 0xd8, 0x9b, 0x15,
	0x4c, 0x30, 0xab, 0x27, 0x55, 0x6f, 0xa3, 0x7a, 0xf3, 0x83, 0xc1, 0xe3, 0x94, 0xa5, 0x4c, 0x89,
	0xbe, 0x3c, 0xe9, 0xff, 0x0d, 0x9e, 0xc6, 0x8c, 0xe7, 0x8c, 0x87, 0x5a, 0xd0, 0x85, 0x96, 0xe0,
	0xb7, 0x6d, 0xb0, 0x73, 0x8a, 0xc5, 0x51, 0xce, 0x4a, 0x2a, 0xc6, 0xd2, 0xdb, 0x0a, 0x41, 0x37,
	0xcf, 0xa8, 0x08, 0x0b, 0x24, 0xb0, 0x6d, 0x3e, 0x33, 0x5f, 0x74, 0x47, 0xa3, 0xcb, 0xa5, 0x6b,
	0xfc, 0x58, 0xba, 0xcf, 0xd3, 0x4c, 0x4c, 0xca, 0xc8, 0x8b, 0x59, 0x5e, 0xd9, 0x54, 0x3f, 0x43,
	0x9e, 0x4c, 0x7d, 0xb1, 0x98, 0x61, 0xee, 0x1d, 0xe3, 0xf8, 0x7a, 0xe9, 0xf6, 0x16, 0x28, 0x27,
	0x87, 0xb0, 0x36, 0x82, 0x41, 0x47, 0x9e, 0x03, 0x09, 0x98, 0x80, 0x87, 0x82, 0x09, 0x44, 0x42,
	0x5e, 0xce, 0x66, 0x64, 0x61, 0xff, 0xa7, 0x18, 0xaf, 0x5b, 0x30, 0xde, 0x52, 0x71, 0xbd, 0x74,
	0x77, 0x35, 0xa3, 0xe9, 0x05, 0x83, 0x07, 0xaa, 0x1c, 0xab, 0xca, 0x8a, 0x00, 0xa0, 0x58, 0x84,
	0x48, 0xdd, 0xce, 0xfe, 0x5f, 0x71, 0x5e, 0xb5, 0xe6, 0xf4, 0x35, 0xe7, 0xc6, 0x09, 0x06, 0x5d,
	0xba, 0x99, 0x99, 0x75, 0x6e, 0x82, 0x27, 0x3a, 0x42, 0x82, 0x09, 0x4e, 0x91, 0xc8, 0x18, 0x0d,
	0xf9, 0x04, 0x15, 0x98, 0xdb, 0xf7, 0x14, 0xf1, 0x5d, 0xeb, 0xe9, 0x39, 0xcd, 0x9b, 0xdd, 0xb1,
	0x85, 0xc1, 0x9e, 0x52, 0x8e, 0x6b, 0x61, 0xac, 0xfa, 0xd6, 0x67, 0xb0, 0xab, 0x3f, 0x21, 0xd9,
	0xc7, 0x32, 0x4b, 0x42, 0xc1, 0xa6, 0x98, 0x72, 0xfb, 0xbe, 0x4a, 0x71, 0xd2, 0xfa, 0xde, 0x83,
	0x66, 0x8a, 0x5b, 0x96, 0x30, 0xe8, 0xab, 0xee, 0x89, 0x6a, 0xbe, 0x57, 0x3d, 0xeb, 0x0c, 0x3c,
	0x2a, 0xf0, 0x19, 0x2a, 0x12, 0x1e, 0x46, 0x88, 0x20, 0x1a, 0x63, 0x7b, 0x4b, 0x91, 0x4f, 0x5b,
	0x93, 0xf7, 0x9b, 0xe4, 0x3f, 0x4c, 0x61, 0xb0, 0x53, 0x75, 0x46, 0xba, 0x61, 0x7d, 0x01, 0xfd,
	0x92, 0x46, 0x8c, 0x26, 0x19, 0x4d, 0x6b, 0xf4, 0x76, 0xeb, 0xd1, 0x6b, 0xf4, 0xad, 0xd1, 0xdf,
	0xb1, 0x85, 0x41, 0xaf, 0xee, 0x6d, 0xf0, 0xf2, 0x01, 0x14, 0x98, 0xe3, 0x62, 0x8e, 0x43, 0x14,
	0xc7, 0xf2, 0x51, 0xd4, 0x29, 0x3a, 0xff, 0x96, 0xe2, 0x2f, 0xb6, 0x30, 0xd8, 0xab, 0x94, 0x23,
	0x2d, 0x54, 0x51, 0x0e, 0x3b, 0xe7, 0x17, 0xae, 0xf1, 0xeb, 0xc2, 0x35, 0x47, 0x6f, 0x2e, 0x57,
	0x8e, 0x79, 0xb5, 0x72, 0xcc, 0x9f, 0x2b, 0xc7, 0xfc, 0xba, 0x76, 0x8c, 0xab, 0xb5, 0x63, 0x7c,
	0x5f, 0x3b, 0xc6, 0x87, 0x61, 0x23, 0x84, 0xdc, 0x1f, 0x43, 0xb9, 0x51, 0x32, 0x46, 0x55, 0xe1,
	0x7f, 0xba, 0x59, 0x36, 0x2a, 0x4f, 0xb4, 0xa5, 0xf6, 0xc4, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbe, 0xa0, 0x91, 0xa8, 0x8a, 0x04, 0x00, 0x00,
}

func (this *NetAmountState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NetAmountState)
	if !ok {
		that2, ok := that.(NetAmountState)
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
	if !this.MintRate.Equal(that1.MintRate) {
		return false
	}
	if !this.TotalSupply.Equal(that1.TotalSupply) {
		return false
	}
	if !this.NetAmount.Equal(that1.NetAmount) {
		return false
	}
	if !this.TotalDelegationShares.Equal(that1.TotalDelegationShares) {
		return false
	}
	if !this.TotalLiquidTokens.Equal(that1.TotalLiquidTokens) {
		return false
	}
	if !this.RewardsBalance.Equal(that1.RewardsBalance) {
		return false
	}
	if !this.UnbondingBalance.Equal(that1.UnbondingBalance) {
		return false
	}
	if !this.ReserveAccountBalance.Equal(that1.ReserveAccountBalance) {
		return false
	}
	return true
}
func (m *NetAmountState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetAmountState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetAmountState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ReserveAccountBalance.Size()
		i -= size
		if _, err := m.ReserveAccountBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.UnbondingBalance.Size()
		i -= size
		if _, err := m.UnbondingBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.RewardsBalance.Size()
		i -= size
		if _, err := m.RewardsBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.TotalLiquidTokens.Size()
		i -= size
		if _, err := m.TotalLiquidTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TotalDelegationShares.Size()
		i -= size
		if _, err := m.TotalDelegationShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.NetAmount.Size()
		i -= size
		if _, err := m.NetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MintRate.Size()
		i -= size
		if _, err := m.MintRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetAmountState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MintRate.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.TotalSupply.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.NetAmount.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.TotalDelegationShares.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.TotalLiquidTokens.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.RewardsBalance.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.UnbondingBalance.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.ReserveAccountBalance.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetAmountState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: NetAmountState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetAmountState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDelegationShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalDelegationShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalLiquidTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalLiquidTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardsBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnbondingBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAccountBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)