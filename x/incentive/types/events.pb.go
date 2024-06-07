// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/incentive/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

type EventNewReward struct {
	DelegatorAddress string        `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorAddress string        `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Reward           types.DecCoin `protobuf:"bytes,3,opt,name=reward,proto3" json:"reward"`
}

func (m *EventNewReward) Reset()         { *m = EventNewReward{} }
func (m *EventNewReward) String() string { return proto.CompactTextString(m) }
func (*EventNewReward) ProtoMessage()    {}
func (*EventNewReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae90c848c5ca7f79, []int{0}
}
func (m *EventNewReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNewReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNewReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNewReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewReward.Merge(m, src)
}
func (m *EventNewReward) XXX_Size() int {
	return m.Size()
}
func (m *EventNewReward) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewReward.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewReward proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventNewReward)(nil), "ollo.incentive.v1.EventNewReward")
}

func init() { proto.RegisterFile("ollo/incentive/v1/events.proto", fileDescriptor_ae90c848c5ca7f79) }

var fileDescriptor_ae90c848c5ca7f79 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x63, 0x40, 0x95, 0x08, 0x12, 0xa2, 0x15, 0x43, 0xa9, 0x90, 0xa9, 0x98, 0x2a, 0x21,
	0x6c, 0x05, 0x36, 0x36, 0x0a, 0x6c, 0x88, 0x21, 0x23, 0x4b, 0xe5, 0x24, 0xa7, 0x60, 0x29, 0xf5,
	0x55, 0xb1, 0x71, 0xe1, 0x2d, 0x78, 0x0a, 0x9e, 0x25, 0x63, 0x47, 0x26, 0x04, 0xc9, 0x8b, 0x20,
	0xc7, 0x21, 0x82, 0xcd, 0x77, 0xdf, 0x77, 0xbe, 0x9f, 0x90, 0x62, 0x51, 0x20, 0x97, 0x2a, 0x05,
	0x65, 0xa4, 0x05, 0x6e, 0x23, 0x0e, 0x16, 0x94, 0xd1, 0x6c, 0x55, 0xa2, 0xc1, 0xd1, 0xd0, 0x71,
	0xd6, 0x73, 0x66, 0xa3, 0xc9, 0x51, 0x8a, 0x7a, 0x89, 0x7a, 0xd1, 0x0a, 0xdc, 0x07, 0xde, 0x9e,
	0x50, 0x1f, 0xf1, 0x44, 0x68, 0xf7, 0x55, 0x02, 0x46, 0x44, 0x3c, 0x45, 0xa9, 0x3a, 0x7e, 0x98,
	0x63, 0x8e, 0xbe, 0xce, 0xbd, 0x7c, 0xf6, 0xf4, 0x9d, 0x84, 0xfb, 0x77, 0xae, 0xe9, 0x03, 0xac,
	0x63, 0x58, 0x8b, 0x32, 0x1b, 0x9d, 0x85, 0xc3, 0x0c, 0x0a, 0xc8, 0x85, 0xc1, 0x72, 0x21, 0xb2,
	0xac, 0x04, 0xad, 0xc7, 0x64, 0x4a, 0x66, 0xbb, 0xf1, 0x41, 0x0f, 0xae, 0x7d, 0xde, 0xc9, 0x56,
	0x14, 0x32, 0xfb, 0x27, 0x6f, 0x79, 0xb9, 0x07, 0xbf, 0xf2, 0x55, 0x38, 0x28, 0xdb, 0x1e, 0xe3,
	0xed, 0x29, 0x99, 0xed, 0x5d, 0x1c, 0xb3, 0x6e, 0x03, 0x37, 0x33, 0xeb, 0x66, 0x66, 0xb7, 0x90,
	0xde, 0xa0, 0x54, 0xf3, 0x9d, 0xea, 0xf3, 0x24, 0x88, 0xbb, 0x8a, 0xf9, 0x7d, 0xf5, 0x4d, 0x83,
	0xaa, 0xa6, 0x64, 0x53, 0x53, 0xf2, 0x55, 0x53, 0xf2, 0xd6, 0xd0, 0x60, 0xd3, 0xd0, 0xe0, 0xa3,
	0xa1, 0xc1, 0x23, 0xcb, 0xa5, 0x79, 0x7a, 0x4e, 0x58, 0x8a, 0x4b, 0xee, 0xae, 0x76, 0xae, 0x8d,
	0x30, 0x12, 0x55, 0x1b, 0xf0, 0x97, 0x3f, 0x47, 0x36, 0xaf, 0x2b, 0xd0, 0xc9, 0xa0, 0xdd, 0xfe,
	0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x00, 0x22, 0x13, 0x55, 0x83, 0x01, 0x00, 0x00,
}

func (m *EventNewReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNewReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNewReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventNewReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Reward.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventNewReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventNewReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNewReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
