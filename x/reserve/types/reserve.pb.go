// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/reserve/v1/reserve.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type DenomWhitelist struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	// repeat
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty" yaml:"addresses"`
}

func (m *DenomWhitelist) Reset()         { *m = DenomWhitelist{} }
func (m *DenomWhitelist) String() string { return proto.CompactTextString(m) }
func (*DenomWhitelist) ProtoMessage()    {}
func (*DenomWhitelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b0325b306dd8356, []int{0}
}
func (m *DenomWhitelist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomWhitelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomWhitelist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomWhitelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomWhitelist.Merge(m, src)
}
func (m *DenomWhitelist) XXX_Size() int {
	return m.Size()
}
func (m *DenomWhitelist) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomWhitelist.DiscardUnknown(m)
}

var xxx_messageInfo_DenomWhitelist proto.InternalMessageInfo

func (m *DenomWhitelist) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DenomWhitelist) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*DenomWhitelist)(nil), "ollo.reserve.v1.DenomWhitelist")
}

func init() { proto.RegisterFile("ollo/reserve/v1/reserve.proto", fileDescriptor_8b0325b306dd8356) }

var fileDescriptor_8b0325b306dd8356 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcf, 0xc9, 0xc9,
	0xd7, 0x2f, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b, 0xd5, 0x2f, 0x33, 0x84, 0x31, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x41, 0xd2, 0x7a, 0x30, 0xb1, 0x32, 0x43, 0x29, 0x91, 0xf4, 0xfc,
	0xf4, 0x7c, 0xb0, 0x9c, 0x3e, 0x88, 0x05, 0x51, 0x26, 0x25, 0x99, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f,
	0x1c, 0x0f, 0x91, 0x80, 0x70, 0x20, 0x52, 0x4a, 0x7d, 0x8c, 0x5c, 0x7c, 0x2e, 0xa9, 0x79, 0xf9,
	0xb9, 0xe1, 0x19, 0x99, 0x25, 0xa9, 0x39, 0x99, 0xc5, 0x25, 0x42, 0x6a, 0x5c, 0xac, 0x29, 0x20,
	0x11, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x81, 0x4f, 0xf7, 0xe4, 0x79, 0x2a, 0x13, 0x73,
	0x73, 0xac, 0x94, 0xc0, 0xc2, 0x4a, 0x41, 0x10, 0x69, 0x21, 0x3f, 0x2e, 0xce, 0xc4, 0x94, 0x94,
	0xa2, 0xd4, 0xe2, 0xe2, 0xd4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x4e, 0x27, 0x83, 0x13, 0xf7,
	0xe4, 0x19, 0x3e, 0xdd, 0x93, 0x17, 0x80, 0xa8, 0x87, 0x4b, 0x2b, 0x5d, 0xda, 0xa2, 0x2b, 0x02,
	0xb5, 0xd7, 0x11, 0x22, 0x18, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x1e, 0x84, 0x30, 0xc2, 0x8a, 0xe5,
	0xc5, 0x02, 0x79, 0x46, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88,
	0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xf9, 0x5b, 0xb7,
	0xb8, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xcc, 0xd1, 0xaf, 0x80, 0x87, 0x52, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x7f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb8, 0x53,
	0xa9, 0x42, 0x01, 0x00, 0x00,
}

func (this *DenomWhitelist) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DenomWhitelist)
	if !ok {
		that2, ok := that.(DenomWhitelist)
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
	if this.Denom != that1.Denom {
		return false
	}
	if len(this.Addresses) != len(that1.Addresses) {
		return false
	}
	for i := range this.Addresses {
		if this.Addresses[i] != that1.Addresses[i] {
			return false
		}
	}
	return true
}
func (m *DenomWhitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomWhitelist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomWhitelist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintReserve(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintReserve(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReserve(dAtA []byte, offset int, v uint64) int {
	offset -= sovReserve(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomWhitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovReserve(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovReserve(uint64(l))
		}
	}
	return n
}

func sovReserve(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReserve(x uint64) (n int) {
	return sovReserve(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomWhitelist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReserve
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
			return fmt.Errorf("proto: DenomWhitelist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomWhitelist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReserve
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
				return ErrInvalidLengthReserve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReserve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReserve
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
				return ErrInvalidLengthReserve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReserve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReserve(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReserve
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
func skipReserve(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReserve
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
					return 0, ErrIntOverflowReserve
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
					return 0, ErrIntOverflowReserve
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
				return 0, ErrInvalidLengthReserve
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReserve
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReserve
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReserve        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReserve          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReserve = fmt.Errorf("proto: unexpected end of group")
)
