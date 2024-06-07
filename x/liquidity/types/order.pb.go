// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/liquidity/v1/order.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// OrderType enumerates order types.
type OrderType int32

const (
	// ORDER_TYPE_UNSPECIFIED specifies unknown order type.
	OrderTypeUnspecified OrderType = 0
	// ORDER_TYPE_LIMIT specifies limit order type.
	OrderTypeLimit OrderType = 1
	// ORDER_TYPE_MARKET specifies market order type.
	OrderTypeMarket OrderType = 2
	// ORDER_TYPE_MM specifies MM(market making) order type.
	OrderTypeMM OrderType = 3
)

var OrderType_name = map[int32]string{
	0: "ORDER_TYPE_UNSPECIFIED",
	1: "ORDER_TYPE_LIMIT",
	2: "ORDER_TYPE_MARKET",
	3: "ORDER_TYPE_MM",
}

var OrderType_value = map[string]int32{
	"ORDER_TYPE_UNSPECIFIED": 0,
	"ORDER_TYPE_LIMIT":       1,
	"ORDER_TYPE_MARKET":      2,
	"ORDER_TYPE_MM":          3,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{0}
}

// OrderDirection enumerates order directions.
type OrderDirection int32

const (
	// ORDER_DIRECTION_UNSPECIFIED specifies unknown order direction
	OrderDirectionUnspecified OrderDirection = 0
	// ORDER_DIRECTION_BUY specifies buy(swap quote coin to base coin) order direction
	OrderDirectionBuy OrderDirection = 1
	// ORDER_DIRECTION_SELL specifies sell(swap base coin to quote coin) order direction
	OrderDirectionSell OrderDirection = 2
)

var OrderDirection_name = map[int32]string{
	0: "ORDER_DIRECTION_UNSPECIFIED",
	1: "ORDER_DIRECTION_BUY",
	2: "ORDER_DIRECTION_SELL",
}

var OrderDirection_value = map[string]int32{
	"ORDER_DIRECTION_UNSPECIFIED": 0,
	"ORDER_DIRECTION_BUY":         1,
	"ORDER_DIRECTION_SELL":        2,
}

func (x OrderDirection) String() string {
	return proto.EnumName(OrderDirection_name, int32(x))
}

func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{1}
}

// RequestStatus enumerates request statuses.
type RequestStatus int32

const (
	// REQUEST_STATUS_UNSPECIFIED specifies unknown request status
	RequestStatusUnspecified RequestStatus = 0
	// REQUEST_STATUS_NOT_EXECUTED indicates the request is not executed yet
	RequestStatusNotExecuted RequestStatus = 1
	// REQUEST_STATUS_SUCCEEDED indicates the request has been succeeded
	RequestStatusSucceeded RequestStatus = 2
	// REQUEST_STATUS_FAILED indicates the request is failed
	RequestStatusFailed RequestStatus = 3
)

var RequestStatus_name = map[int32]string{
	0: "REQUEST_STATUS_UNSPECIFIED",
	1: "REQUEST_STATUS_NOT_EXECUTED",
	2: "REQUEST_STATUS_SUCCEEDED",
	3: "REQUEST_STATUS_FAILED",
}

var RequestStatus_value = map[string]int32{
	"REQUEST_STATUS_UNSPECIFIED":  0,
	"REQUEST_STATUS_NOT_EXECUTED": 1,
	"REQUEST_STATUS_SUCCEEDED":    2,
	"REQUEST_STATUS_FAILED":       3,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{2}
}

// OrderStatus enumerates order statuses.
type OrderStatus int32

const (
	// ORDER_STATUS_UNSPECIFIED specifies unknown order status
	OrderStatusUnspecified OrderStatus = 0
	// ORDER_STATUS_NOT_EXECUTED indicates the order has not been executed yet
	OrderStatusNotExecuted OrderStatus = 1
	// ORDER_STATUS_NOT_MATCHED indicates the order has been executed but has no match
	OrderStatusNotMatched OrderStatus = 2
	// ORDER_STATUS_PARTIALLY_MATCHED indicates the order has been partially matched
	OrderStatusPartiallyMatched OrderStatus = 3
	// ORDER_STATUS_COMPLETED indicates the order has been fully matched and completed
	OrderStatusCompleted OrderStatus = 4
	// ORDER_STATUS_CANCELED indicates the order has been canceled
	OrderStatusCanceled OrderStatus = 5
	// ORDER_STATUS_EXPIRED indicates the order has been expired
	OrderStatusExpired OrderStatus = 6
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_STATUS_UNSPECIFIED",
	1: "ORDER_STATUS_NOT_EXECUTED",
	2: "ORDER_STATUS_NOT_MATCHED",
	3: "ORDER_STATUS_PARTIALLY_MATCHED",
	4: "ORDER_STATUS_COMPLETED",
	5: "ORDER_STATUS_CANCELED",
	6: "ORDER_STATUS_EXPIRED",
}

var OrderStatus_value = map[string]int32{
	"ORDER_STATUS_UNSPECIFIED":       0,
	"ORDER_STATUS_NOT_EXECUTED":      1,
	"ORDER_STATUS_NOT_MATCHED":       2,
	"ORDER_STATUS_PARTIALLY_MATCHED": 3,
	"ORDER_STATUS_COMPLETED":         4,
	"ORDER_STATUS_CANCELED":          5,
	"ORDER_STATUS_EXPIRED":           6,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{3}
}

// Order defines an order.
type Order struct {
	// type specifies the typo of the order
	Type OrderType `protobuf:"varint,1,opt,name=type,proto3,enum=ollo.liquidity.v1.OrderType" json:"type,omitempty"`
	// id specifies the id of the order
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// pair_id specifies the pair id
	PairId uint64 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	// msg_height specifies the block height when the order is stored for the batch execution
	MsgHeight int64 `protobuf:"varint,4,opt,name=msg_height,json=msgHeight,proto3" json:"msg_height,omitempty"`
	// orderer specifies the bech32-encoded address that makes an order
	Orderer string `protobuf:"bytes,5,opt,name=orderer,proto3" json:"orderer,omitempty" yaml:"orderer"`
	// direction specifies the order direction; either buy or sell
	Direction OrderDirection `protobuf:"varint,6,opt,name=direction,proto3,enum=ollo.liquidity.v1.OrderDirection" json:"direction,omitempty"`
	OfferCoin types.Coin     `protobuf:"bytes,7,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
	// remaining_offer_coin specifies the remaining offer coin
	RemainingOfferCoin types.Coin `protobuf:"bytes,8,opt,name=remaining_offer_coin,json=remainingOfferCoin,proto3" json:"remaining_offer_coin"`
	// received_coin specifies the received coin after the swap
	ReceivedCoin types.Coin `protobuf:"bytes,9,opt,name=received_coin,json=receivedCoin,proto3" json:"received_coin"`
	// price specifies the price that an orderer is willing to swap
	Price      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Amount     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	OpenAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=open_amount,json=openAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"open_amount"`
	// batch_id specifies the pair's batch id when the request is stored
	BatchId  uint64      `protobuf:"varint,13,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	ExpireAt time.Time   `protobuf:"bytes,14,opt,name=expire_at,json=expireAt,proto3,stdtime" json:"expire_at"`
	Status   OrderStatus `protobuf:"varint,15,opt,name=status,proto3,enum=ollo.liquidity.v1.OrderStatus" json:"status,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_71b117091565b732, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.liquidity.v1.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("ollo.liquidity.v1.OrderDirection", OrderDirection_name, OrderDirection_value)
	proto.RegisterEnum("ollo.liquidity.v1.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterEnum("ollo.liquidity.v1.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*Order)(nil), "ollo.liquidity.v1.Order")
}

func init() { proto.RegisterFile("ollo/liquidity/v1/order.proto", fileDescriptor_71b117091565b732) }

var fileDescriptor_71b117091565b732 = []byte{
	// 1096 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xcd, 0x4e, 0xe3, 0x56,
	0x14, 0xc7, 0xe3, 0x24, 0x04, 0x72, 0x19, 0x32, 0xe1, 0xf2, 0x31, 0x8e, 0x19, 0x1c, 0x17, 0x55,
	0x55, 0x84, 0x84, 0x33, 0xd0, 0xaa, 0x5f, 0x6a, 0xa9, 0xf2, 0x71, 0xd1, 0x58, 0x4d, 0x48, 0xc6,
	0x71, 0x24, 0xe8, 0x26, 0x72, 0xec, 0x4b, 0xb8, 0x1a, 0x27, 0xce, 0xd8, 0x37, 0x88, 0xec, 0xba,
	0xac, 0xbc, 0x9a, 0x17, 0xf0, 0xa6, 0xdd, 0xf4, 0x01, 0xba, 0xe8, 0xb6, 0x3b, 0x96, 0xa8, 0xab,
	0xaa, 0x0b, 0x3a, 0xc0, 0x1b, 0xf4, 0x09, 0x2a, 0x7f, 0x24, 0x38, 0xe9, 0x8c, 0x34, 0xed, 0x0a,
	0xae, 0xcf, 0xff, 0xf7, 0xf7, 0x39, 0xc7, 0xe7, 0x5c, 0x05, 0x6c, 0x9b, 0x86, 0x61, 0x16, 0x0d,
	0xf2, 0x6a, 0x44, 0x74, 0x42, 0xc7, 0xc5, 0x8b, 0xfd, 0xa2, 0x69, 0xe9, 0xd8, 0x12, 0x87, 0x96,
	0x49, 0x4d, 0xb8, 0xea, 0x85, 0xc5, 0x69, 0x58, 0xbc, 0xd8, 0xe7, 0xd6, 0x7b, 0x66, 0xcf, 0xf4,
	0xa3, 0x45, 0xef, 0xbf, 0x40, 0xc8, 0xe5, 0x34, 0xd3, 0xee, 0x9b, 0x76, 0x27, 0x08, 0x04, 0x87,
	0x30, 0xc4, 0x07, 0xa7, 0x62, 0x57, 0xb5, 0x71, 0xf1, 0x62, 0xbf, 0x8b, 0xa9, 0xba, 0x5f, 0xd4,
	0x4c, 0x32, 0x08, 0xe3, 0xf9, 0x9e, 0x69, 0xf6, 0x0c, 0x5c, 0xf4, 0x4f, 0xdd, 0xd1, 0x59, 0x91,
	0x92, 0x3e, 0xb6, 0xa9, 0xda, 0x1f, 0x4e, 0x0c, 0xe6, 0x05, 0xfa, 0xc8, 0x52, 0x29, 0x31, 0x43,
	0x83, 0x9d, 0xeb, 0x14, 0x58, 0x68, 0x78, 0x49, 0xc3, 0x67, 0x20, 0x49, 0xc7, 0x43, 0xcc, 0x32,
	0x02, 0x53, 0xc8, 0x1c, 0x3c, 0x15, 0xff, 0x95, 0xbd, 0xe8, 0xeb, 0x94, 0xf1, 0x10, 0xcb, 0xbe,
	0x12, 0x66, 0x40, 0x9c, 0xe8, 0x6c, 0x5c, 0x60, 0x0a, 0x49, 0x39, 0x4e, 0x74, 0xf8, 0x04, 0x2c,
	0x0e, 0x55, 0x62, 0x75, 0x88, 0xce, 0x26, 0xfc, 0x87, 0x29, 0xef, 0x28, 0xe9, 0x70, 0x1b, 0x80,
	0xbe, 0xdd, 0xeb, 0x9c, 0x63, 0xd2, 0x3b, 0xa7, 0x6c, 0x52, 0x60, 0x0a, 0x09, 0x39, 0xdd, 0xb7,
	0x7b, 0xcf, 0xfd, 0x07, 0xf0, 0x10, 0x2c, 0xfa, 0x7d, 0xc3, 0x16, 0xbb, 0x20, 0x30, 0x85, 0x74,
	0xf9, 0xc3, 0xbf, 0x6f, 0xf2, 0x99, 0xb1, 0xda, 0x37, 0xbe, 0xdc, 0x09, 0x03, 0x3b, 0xbf, 0xff,
	0xb2, 0x97, 0x09, 0x3b, 0x53, 0xd2, 0x75, 0x0b, 0xdb, 0xb6, 0x3c, 0x81, 0xe0, 0x37, 0x20, 0xad,
	0x13, 0x0b, 0x6b, 0x5e, 0x59, 0x6c, 0xca, 0x4f, 0xff, 0x83, 0x77, 0xa5, 0x5f, 0x9d, 0x08, 0xe5,
	0x07, 0x06, 0x1e, 0x02, 0x60, 0x9e, 0x9d, 0x61, 0xab, 0xe3, 0x75, 0x96, 0x5d, 0x14, 0x98, 0xc2,
	0xf2, 0x41, 0x4e, 0x0c, 0x5f, 0xe7, 0xb5, 0x5e, 0x0c, 0x5b, 0x2f, 0x56, 0x4c, 0x32, 0x28, 0x27,
	0xaf, 0x6e, 0xf2, 0x31, 0x39, 0xed, 0x23, 0xde, 0x03, 0xf8, 0x02, 0xac, 0x5b, 0xb8, 0xaf, 0x92,
	0x01, 0x19, 0xf4, 0x3a, 0x11, 0xa7, 0xa5, 0xf7, 0x73, 0x82, 0x53, 0xb8, 0x31, 0xb5, 0xac, 0x82,
	0x15, 0x0b, 0x6b, 0x98, 0x5c, 0x60, 0x3d, 0xf0, 0x4a, 0xbf, 0x9f, 0xd7, 0xa3, 0x09, 0x15, 0xba,
	0x2c, 0x0c, 0x2d, 0xa2, 0x61, 0x16, 0xf8, 0x7d, 0x15, 0x3d, 0xc9, 0x9f, 0x37, 0xf9, 0x8f, 0x7a,
	0x84, 0x9e, 0x8f, 0xba, 0xa2, 0x66, 0xf6, 0xc3, 0x71, 0x0b, 0xff, 0xec, 0xd9, 0xfa, 0xcb, 0xa2,
	0xf7, 0x6d, 0x6d, 0xb1, 0x8a, 0x35, 0x39, 0x80, 0xe1, 0x11, 0x48, 0xa9, 0x7d, 0x73, 0x34, 0xa0,
	0xec, 0xf2, 0x7f, 0xb6, 0x91, 0x06, 0x54, 0x0e, 0x69, 0xd8, 0x00, 0xcb, 0xe6, 0x10, 0x0f, 0x3a,
	0xa1, 0xd9, 0xa3, 0xff, 0x65, 0x06, 0x3c, 0x8b, 0x52, 0x60, 0x98, 0x03, 0x4b, 0x5d, 0x95, 0x6a,
	0xe7, 0xde, 0xc4, 0xad, 0xf8, 0x13, 0xb7, 0xe8, 0x9f, 0x25, 0x1d, 0x96, 0x40, 0x1a, 0x5f, 0x0e,
	0x89, 0x85, 0x3b, 0x2a, 0x65, 0x33, 0x7e, 0xef, 0x38, 0x31, 0xd8, 0x05, 0x71, 0xb2, 0x0b, 0xa2,
	0x32, 0x59, 0x96, 0xf2, 0x92, 0x97, 0xc5, 0xeb, 0xbf, 0xf2, 0x8c, 0xbc, 0x14, 0x60, 0x25, 0x0a,
	0x3f, 0x05, 0x29, 0x9b, 0xaa, 0x74, 0x64, 0xb3, 0x8f, 0xfd, 0x99, 0xe2, 0xdf, 0x35, 0x53, 0x2d,
	0x5f, 0x25, 0x87, 0xea, 0xdd, 0xdf, 0x18, 0x90, 0x9e, 0xae, 0x0a, 0xfc, 0x04, 0x6c, 0x36, 0xe4,
	0x2a, 0x92, 0x3b, 0xca, 0x69, 0x13, 0x75, 0xda, 0xc7, 0xad, 0x26, 0xaa, 0x48, 0x47, 0x12, 0xaa,
	0x66, 0x63, 0x1c, 0xeb, 0xb8, 0xc2, 0xfa, 0x54, 0xda, 0x1e, 0xd8, 0x43, 0xac, 0x91, 0x33, 0x82,
	0x75, 0x58, 0x00, 0xd9, 0x08, 0x55, 0x93, 0xea, 0x92, 0x92, 0x65, 0x38, 0xe8, 0xb8, 0x42, 0x66,
	0xaa, 0xaf, 0x91, 0x3e, 0xa1, 0x70, 0x17, 0xac, 0x46, 0x94, 0xf5, 0x92, 0xfc, 0x2d, 0x52, 0xb2,
	0x71, 0x6e, 0xcd, 0x71, 0x85, 0xc7, 0x53, 0x69, 0x5d, 0xb5, 0x5e, 0x62, 0x0a, 0x77, 0xc0, 0x4a,
	0x54, 0x5b, 0xcf, 0x26, 0xb8, 0xc7, 0x8e, 0x2b, 0x2c, 0x3f, 0xe8, 0xea, 0x5c, 0xf2, 0x87, 0x9f,
	0xf8, 0xd8, 0xee, 0xaf, 0x0c, 0xc8, 0xcc, 0xee, 0x0b, 0x3c, 0x04, 0x5b, 0x01, 0x5c, 0x95, 0x64,
	0x54, 0x51, 0xa4, 0xc6, 0xf1, 0x5c, 0x35, 0xdb, 0x8e, 0x2b, 0xe4, 0x66, 0xa1, 0x68, 0x49, 0x22,
	0x58, 0x9b, 0xe7, 0xcb, 0xed, 0xd3, 0x2c, 0xc3, 0x6d, 0x38, 0xae, 0xb0, 0x3a, 0xcb, 0x95, 0x47,
	0x63, 0xf8, 0x0c, 0xac, 0xcf, 0xeb, 0x5b, 0xa8, 0x56, 0xcb, 0xc6, 0xb9, 0x4d, 0xc7, 0x15, 0xe0,
	0x2c, 0xd0, 0xc2, 0x86, 0x11, 0xa6, 0xfe, 0x7d, 0x1c, 0xac, 0xc8, 0xf8, 0xd5, 0x08, 0xdb, 0x34,
	0xf8, 0x30, 0xf0, 0x2b, 0xc0, 0xc9, 0xe8, 0x45, 0x1b, 0xb5, 0x94, 0x4e, 0x4b, 0x29, 0x29, 0xed,
	0xd6, 0x5c, 0xe2, 0x4f, 0x1d, 0x57, 0x60, 0x67, 0x90, 0x68, 0xde, 0x5f, 0x83, 0xad, 0x39, 0xfa,
	0xb8, 0xa1, 0x74, 0xd0, 0x09, 0xaa, 0xb4, 0x15, 0x54, 0xcd, 0x32, 0x6f, 0xc1, 0x8f, 0x4d, 0x8a,
	0x2e, 0xb1, 0x36, 0xa2, 0x58, 0x87, 0x9f, 0x03, 0x76, 0x0e, 0x6f, 0xb5, 0x2b, 0x15, 0x84, 0xaa,
	0xa8, 0x9a, 0x8d, 0x73, 0x9c, 0xe3, 0x0a, 0x9b, 0x33, 0x6c, 0x6b, 0xa4, 0x69, 0x18, 0xeb, 0x58,
	0x87, 0x07, 0x60, 0x63, 0x8e, 0x3c, 0x2a, 0x49, 0x35, 0x54, 0xcd, 0x26, 0xb8, 0x27, 0x8e, 0x2b,
	0xac, 0xcd, 0x60, 0x47, 0x2a, 0x31, 0xb0, 0x1e, 0xb6, 0xe0, 0xc7, 0x04, 0x58, 0x8e, 0x4c, 0xa6,
	0x97, 0x43, 0xd0, 0xca, 0xb7, 0x96, 0xef, 0xe7, 0x10, 0x91, 0x47, 0x8b, 0xff, 0x02, 0xe4, 0x66,
	0xc8, 0xb9, 0xd2, 0xe7, 0xd1, 0x68, 0xe1, 0x9f, 0xcd, 0xbd, 0xd4, 0x43, 0xeb, 0x25, 0xa5, 0xf2,
	0xdc, 0x2f, 0x3c, 0xe7, 0xb8, 0xc2, 0xc6, 0x2c, 0x59, 0xf7, 0x56, 0x17, 0xeb, 0xb0, 0x02, 0xf8,
	0x19, 0xb0, 0x59, 0x92, 0x15, 0xa9, 0x54, 0xab, 0x9d, 0x4e, 0xf1, 0x04, 0x97, 0x77, 0x5c, 0x61,
	0x2b, 0x82, 0x37, 0x55, 0x8b, 0x12, 0xd5, 0x30, 0xc6, 0x13, 0x93, 0xe9, 0xda, 0x85, 0x26, 0x95,
	0x46, 0xbd, 0x59, 0x43, 0x5e, 0xd6, 0xc9, 0xc8, 0xda, 0x05, 0x70, 0xc5, 0xec, 0x0f, 0x0d, 0x4c,
	0x83, 0x96, 0xcf, 0x52, 0xa5, 0xe3, 0x0a, 0xf2, 0x5a, 0xbe, 0x10, 0xb4, 0x3c, 0x0a, 0xa9, 0x03,
	0x0d, 0x1b, 0x58, 0x7f, 0x98, 0xd3, 0x90, 0x41, 0x27, 0x4d, 0x49, 0x46, 0xd5, 0x6c, 0x2a, 0x32,
	0xa7, 0x01, 0x82, 0xfc, 0x9b, 0x25, 0xfc, 0x48, 0xe5, 0x93, 0xab, 0x5b, 0x3e, 0x76, 0x7d, 0xcb,
	0xc7, 0xde, 0xdc, 0xf2, 0xcc, 0xcf, 0x77, 0x3c, 0x73, 0x75, 0xc7, 0x33, 0xd7, 0x77, 0x3c, 0xf3,
	0xe6, 0x8e, 0x67, 0x5e, 0xdf, 0xf3, 0xb1, 0xeb, 0x7b, 0x3e, 0xf6, 0xc7, 0x3d, 0x1f, 0xfb, 0x4e,
	0x8c, 0x5c, 0x8d, 0xde, 0x15, 0xb4, 0xe7, 0x5d, 0x34, 0xc4, 0x1c, 0xf8, 0x87, 0xe2, 0x65, 0xe4,
	0x17, 0x88, 0x7f, 0x4d, 0x76, 0x53, 0xfe, 0x05, 0xf7, 0xf1, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0xec, 0x7d, 0x64, 0xa0, 0x08, 0x00, 0x00,
}

func (this *Order) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Order)
	if !ok {
		that2, ok := that.(Order)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Order")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Order but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Order but is not nil && this == nil")
	}
	if this.Type != that1.Type {
		return fmt.Errorf("Type this(%v) Not Equal that(%v)", this.Type, that1.Type)
	}
	if this.Id != that1.Id {
		return fmt.Errorf("Id this(%v) Not Equal that(%v)", this.Id, that1.Id)
	}
	if this.PairId != that1.PairId {
		return fmt.Errorf("PairId this(%v) Not Equal that(%v)", this.PairId, that1.PairId)
	}
	if this.MsgHeight != that1.MsgHeight {
		return fmt.Errorf("MsgHeight this(%v) Not Equal that(%v)", this.MsgHeight, that1.MsgHeight)
	}
	if this.Orderer != that1.Orderer {
		return fmt.Errorf("Orderer this(%v) Not Equal that(%v)", this.Orderer, that1.Orderer)
	}
	if this.Direction != that1.Direction {
		return fmt.Errorf("Direction this(%v) Not Equal that(%v)", this.Direction, that1.Direction)
	}
	if !this.OfferCoin.Equal(&that1.OfferCoin) {
		return fmt.Errorf("OfferCoin this(%v) Not Equal that(%v)", this.OfferCoin, that1.OfferCoin)
	}
	if !this.RemainingOfferCoin.Equal(&that1.RemainingOfferCoin) {
		return fmt.Errorf("RemainingOfferCoin this(%v) Not Equal that(%v)", this.RemainingOfferCoin, that1.RemainingOfferCoin)
	}
	if !this.ReceivedCoin.Equal(&that1.ReceivedCoin) {
		return fmt.Errorf("ReceivedCoin this(%v) Not Equal that(%v)", this.ReceivedCoin, that1.ReceivedCoin)
	}
	if !this.Price.Equal(that1.Price) {
		return fmt.Errorf("Price this(%v) Not Equal that(%v)", this.Price, that1.Price)
	}
	if !this.Amount.Equal(that1.Amount) {
		return fmt.Errorf("Amount this(%v) Not Equal that(%v)", this.Amount, that1.Amount)
	}
	if !this.OpenAmount.Equal(that1.OpenAmount) {
		return fmt.Errorf("OpenAmount this(%v) Not Equal that(%v)", this.OpenAmount, that1.OpenAmount)
	}
	if this.BatchId != that1.BatchId {
		return fmt.Errorf("BatchId this(%v) Not Equal that(%v)", this.BatchId, that1.BatchId)
	}
	if !this.ExpireAt.Equal(that1.ExpireAt) {
		return fmt.Errorf("ExpireAt this(%v) Not Equal that(%v)", this.ExpireAt, that1.ExpireAt)
	}
	if this.Status != that1.Status {
		return fmt.Errorf("Status this(%v) Not Equal that(%v)", this.Status, that1.Status)
	}
	return nil
}
func (this *Order) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Order)
	if !ok {
		that2, ok := that.(Order)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.PairId != that1.PairId {
		return false
	}
	if this.MsgHeight != that1.MsgHeight {
		return false
	}
	if this.Orderer != that1.Orderer {
		return false
	}
	if this.Direction != that1.Direction {
		return false
	}
	if !this.OfferCoin.Equal(&that1.OfferCoin) {
		return false
	}
	if !this.RemainingOfferCoin.Equal(&that1.RemainingOfferCoin) {
		return false
	}
	if !this.ReceivedCoin.Equal(&that1.ReceivedCoin) {
		return false
	}
	if !this.Price.Equal(that1.Price) {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if !this.OpenAmount.Equal(that1.OpenAmount) {
		return false
	}
	if this.BatchId != that1.BatchId {
		return false
	}
	if !this.ExpireAt.Equal(that1.ExpireAt) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x78
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpireAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOrder(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x72
	if m.BatchId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.BatchId))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.OpenAmount.Size()
		i -= size
		if _, err := m.OpenAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.ReceivedCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size, err := m.RemainingOfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Direction != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Orderer) > 0 {
		i -= len(m.Orderer)
		copy(dAtA[i:], m.Orderer)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Orderer)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MsgHeight != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.MsgHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.PairId != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovOrder(uint64(m.Type))
	}
	if m.Id != 0 {
		n += 1 + sovOrder(uint64(m.Id))
	}
	if m.PairId != 0 {
		n += 1 + sovOrder(uint64(m.PairId))
	}
	if m.MsgHeight != 0 {
		n += 1 + sovOrder(uint64(m.MsgHeight))
	}
	l = len(m.Orderer)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovOrder(uint64(m.Direction))
	}
	l = m.OfferCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.RemainingOfferCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.ReceivedCoin.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = m.OpenAmount.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.BatchId != 0 {
		n += 1 + sovOrder(uint64(m.BatchId))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt)
	n += 1 + l + sovOrder(uint64(l))
	if m.Status != 0 {
		n += 1 + sovOrder(uint64(m.Status))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= OrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgHeight", wireType)
			}
			m.MsgHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orderer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= OrderDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingOfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingOfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReceivedCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OpenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
			}
			m.BatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpireAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
