// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/exchange/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	MarketCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=market_creation_fee,json=marketCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"market_creation_fee"`
	Fees              Fees                                     `protobuf:"bytes,2,opt,name=fees,proto3" json:"fees"`
	MaxOrderLifespan  time.Duration                            `protobuf:"bytes,3,opt,name=max_order_lifespan,json=maxOrderLifespan,proto3,stdduration" json:"max_order_lifespan"`
	// max_order_price_ratio defines the ratio of the maximum possible order price compared to the market's last price
	MaxOrderPriceRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_order_price_ratio,json=maxOrderPriceRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_order_price_ratio"`
	MaxSwapRoutesLen   uint32                                 `protobuf:"varint,5,opt,name=max_swap_routes_len,json=maxSwapRoutesLen,proto3" json:"max_swap_routes_len,omitempty"`
	MaxNumMMOrders     uint32                                 `protobuf:"varint,6,opt,name=max_num_mm_orders,json=maxNumMmOrders,proto3" json:"max_num_mm_orders,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_194713c21235dedc, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type Fees struct {
	DefaultMakerFeeRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=default_maker_fee_rate,json=defaultMakerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"default_maker_fee_rate"`
	DefaultTakerFeeRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=default_taker_fee_rate,json=defaultTakerFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"default_taker_fee_rate"`
	DefaultOrderSourceFeeRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=default_order_source_fee_ratio,json=defaultOrderSourceFeeRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"default_order_source_fee_ratio"`
}

func (m *Fees) Reset()         { *m = Fees{} }
func (m *Fees) String() string { return proto.CompactTextString(m) }
func (*Fees) ProtoMessage()    {}
func (*Fees) Descriptor() ([]byte, []int) {
	return fileDescriptor_194713c21235dedc, []int{1}
}
func (m *Fees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fees.Merge(m, src)
}
func (m *Fees) XXX_Size() int {
	return m.Size()
}
func (m *Fees) XXX_DiscardUnknown() {
	xxx_messageInfo_Fees.DiscardUnknown(m)
}

var xxx_messageInfo_Fees proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "ollo.exchange.v1beta1.Params")
	proto.RegisterType((*Fees)(nil), "ollo.exchange.v1beta1.Fees")
}

func init() {
	proto.RegisterFile("ollo/exchange/v1beta1/params.proto", fileDescriptor_194713c21235dedc)
}

var fileDescriptor_194713c21235dedc = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0x24, 0x5f, 0xf4, 0xe1, 0x8a, 0x8a, 0xba, 0x80, 0x9c, 0x2c, 0xec, 0xa8, 0x8b,
	0x2a, 0x9b, 0x8c, 0x69, 0x11, 0x12, 0x2c, 0xd8, 0xb8, 0x55, 0x56, 0x0d, 0x14, 0x17, 0xb1, 0x60,
	0x63, 0x4d, 0x26, 0x37, 0xae, 0x95, 0x8c, 0xc7, 0x9a, 0x19, 0x37, 0x46, 0xbc, 0x04, 0x4b, 0x9e,
	0x81, 0xe7, 0x60, 0x91, 0x65, 0x97, 0x88, 0x45, 0x0a, 0x09, 0x0f, 0x82, 0x66, 0x6c, 0x93, 0xb0,
	0x40, 0xe2, 0xcf, 0xca, 0x1e, 0xcf, 0x3d, 0xe7, 0x77, 0xef, 0x99, 0x91, 0xcd, 0x43, 0xc2, 0x41,
	0x10, 0x48, 0xa4, 0x07, 0x39, 0xb9, 0xc4, 0x49, 0x04, 0xde, 0xd5, 0xd1, 0x08, 0x24, 0x3e, 0xf2,
	0x52, 0xcc, 0x31, 0x15, 0x28, 0xe5, 0x4c, 0x32, 0xab, 0x5d, 0xd5, 0xa1, 0xaa, 0x0e, 0x95, 0x75,
	0x9d, 0xbb, 0x11, 0x8b, 0x98, 0xae, 0xf2, 0xd4, 0x5b, 0x21, 0xe8, 0x38, 0x11, 0x63, 0xd1, 0x0c,
	0x3c, 0xbd, 0x1a, 0x65, 0x13, 0x6f, 0x9c, 0x71, 0x2c, 0x63, 0x96, 0x54, 0xfb, 0x84, 0x09, 0xca,
	0x84, 0x37, 0xc2, 0x62, 0x83, 0x24, 0x2c, 0x2e, 0xf7, 0x0f, 0xbe, 0x35, 0xcc, 0xd6, 0xb9, 0xee,
	0xc0, 0x7a, 0x6b, 0xee, 0x53, 0xcc, 0xa7, 0x20, 0x43, 0xc2, 0x41, 0x7b, 0x84, 0x13, 0x00, 0xdb,
	0xe8, 0x36, 0x7a, 0x3b, 0xc7, 0x6d, 0x54, 0x18, 0x21, 0x65, 0x54, 0xf5, 0x84, 0x4e, 0x58, 0x9c,
	0xf8, 0x0f, 0x16, 0x4b, 0xb7, 0xf6, 0xe1, 0xc6, 0xed, 0x45, 0xb1, 0xbc, 0xcc, 0x46, 0x88, 0x30,
	0xea, 0x95, 0xd4, 0xe2, 0xd1, 0x17, 0xe3, 0xa9, 0x27, 0xdf, 0xa4, 0x20, 0xb4, 0x40, 0x04, 0x7b,
	0x05, 0xe7, 0xa4, 0xc4, 0x0c, 0x00, 0xac, 0x27, 0x66, 0x73, 0x02, 0x20, 0xec, 0x7a, 0xd7, 0xe8,
	0xed, 0x1c, 0xbb, 0xe8, 0x97, 0x39, 0xa0, 0x01, 0x80, 0xf0, 0x9b, 0x8a, 0x19, 0x68, 0x89, 0xf5,
	0xc2, 0xb4, 0x28, 0xce, 0x43, 0xc6, 0xc7, 0xc0, 0xc3, 0x59, 0x3c, 0x01, 0x91, 0xe2, 0xc4, 0x6e,
	0x68, 0xa3, 0x36, 0x2a, 0xf2, 0x41, 0x55, 0x3e, 0xe8, 0xb4, 0xcc, 0xc7, 0xff, 0x5f, 0x59, 0xbc,
	0xbf, 0x71, 0x8d, 0xe0, 0x0e, 0xc5, 0xf9, 0x73, 0xa5, 0x3e, 0x2b, 0xc5, 0x16, 0x36, 0xef, 0x6d,
	0x2c, 0x53, 0x1e, 0x13, 0x08, 0xb5, 0xca, 0x6e, 0x76, 0x8d, 0xde, 0x2d, 0x1f, 0x29, 0xe9, 0xe7,
	0xa5, 0x7b, 0xf8, 0x1b, 0x13, 0x9f, 0x02, 0x09, 0xac, 0x0a, 0x70, 0xae, 0xac, 0x02, 0xe5, 0x64,
	0xf5, 0x55, 0xda, 0x79, 0x28, 0xe6, 0x38, 0x0d, 0x39, 0xcb, 0x24, 0x88, 0x70, 0x06, 0x89, 0xfd,
	0x5f, 0xd7, 0xe8, 0xdd, 0xd6, 0x1d, 0x5d, 0xcc, 0x71, 0x1a, 0xe8, 0x8d, 0x33, 0x48, 0xac, 0xa7,
	0xe6, 0x9e, 0x2a, 0x4f, 0x32, 0x1a, 0x52, 0x5a, 0x34, 0x26, 0xec, 0x96, 0x2a, 0xf6, 0xad, 0xd5,
	0xd2, 0xdd, 0x1d, 0xe2, 0xfc, 0x59, 0x46, 0x87, 0x43, 0x8d, 0x11, 0xc1, 0x2e, 0x2d, 0xd6, 0xb4,
	0x58, 0x1f, 0x7c, 0xac, 0x9b, 0x4d, 0x15, 0x9c, 0x45, 0xcc, 0xfb, 0x63, 0x98, 0xe0, 0x6c, 0x26,
	0x43, 0x8a, 0xa7, 0xc0, 0xd5, 0x11, 0xab, 0xd9, 0xd4, 0x39, 0xff, 0xcd, 0x68, 0xfb, 0xa5, 0xdb,
	0x50, 0x99, 0x0d, 0x40, 0x0d, 0x07, 0xdb, 0x10, 0xf9, 0x33, 0xa4, 0xfe, 0x4f, 0x90, 0x97, 0xdb,
	0x10, 0x6e, 0x3a, 0x15, 0xa4, 0x38, 0x27, 0xc1, 0x32, 0x4e, 0xa0, 0x62, 0xc5, 0x4c, 0x5f, 0x81,
	0x3f, 0x87, 0x75, 0x4a, 0x57, 0x9d, 0xdc, 0x85, 0xf6, 0x2c, 0x90, 0x31, 0xf3, 0x5f, 0x2d, 0xbe,
	0x3a, 0xb5, 0xc5, 0xca, 0x31, 0xae, 0x57, 0x8e, 0xf1, 0x65, 0xe5, 0x18, 0xef, 0xd6, 0x4e, 0xed,
	0x7a, 0xed, 0xd4, 0x3e, 0xad, 0x9d, 0xda, 0xeb, 0xc7, 0xdb, 0x84, 0xf2, 0xfe, 0xf6, 0x13, 0x90,
	0x73, 0xc6, 0xa7, 0x3f, 0x3e, 0x78, 0x57, 0x8f, 0xbc, 0x7c, 0xf3, 0x17, 0xd0, 0xdc, 0x51, 0x4b,
	0x5f, 0xcf, 0x87, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x93, 0xdf, 0xf0, 0x27, 0x04, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxNumMMOrders != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxNumMMOrders))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxSwapRoutesLen != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSwapRoutesLen))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MaxOrderPriceRatio.Size()
		i -= size
		if _, err := m.MaxOrderPriceRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxOrderLifespan, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Fees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketCreationFee) > 0 {
		for iNdEx := len(m.MarketCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MarketCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Fees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.DefaultOrderSourceFeeRatio.Size()
		i -= size
		if _, err := m.DefaultOrderSourceFeeRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DefaultTakerFeeRate.Size()
		i -= size
		if _, err := m.DefaultTakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.DefaultMakerFeeRate.Size()
		i -= size
		if _, err := m.DefaultMakerFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MarketCreationFee) > 0 {
		for _, e := range m.MarketCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.Fees.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan)
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxOrderPriceRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxSwapRoutesLen != 0 {
		n += 1 + sovParams(uint64(m.MaxSwapRoutesLen))
	}
	if m.MaxNumMMOrders != 0 {
		n += 1 + sovParams(uint64(m.MaxNumMMOrders))
	}
	return n
}

func (m *Fees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DefaultMakerFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DefaultTakerFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DefaultOrderSourceFeeRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketCreationFee = append(m.MarketCreationFee, types.Coin{})
			if err := m.MarketCreationFee[len(m.MarketCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderLifespan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxOrderLifespan, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderPriceRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxOrderPriceRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSwapRoutesLen", wireType)
			}
			m.MaxSwapRoutesLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSwapRoutesLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumMMOrders", wireType)
			}
			m.MaxNumMMOrders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumMMOrders |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Fees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Fees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefaultMakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultTakerFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefaultTakerFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultOrderSourceFeeRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DefaultOrderSourceFeeRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
