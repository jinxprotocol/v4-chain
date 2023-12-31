// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/indexer/off_chain_updates/off_chain_updates.proto

package off_chain_updates

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	v1 "github.com/jinxprotocol/v4-chain/protocol/indexer/protocol/v1"
	shared "github.com/jinxprotocol/v4-chain/protocol/indexer/shared"
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

// OrderPlacementStatus is an enum for the resulting status after an order is
// placed.
type OrderPlaceV1_OrderPlacementStatus int32

const (
	// Default value, this is invalid and unused.
	OrderPlaceV1_ORDER_PLACEMENT_STATUS_UNSPECIFIED OrderPlaceV1_OrderPlacementStatus = 0
	// A best effort opened order is one that has only been confirmed to be
	// placed on the jInX node sending the off-chain update message.
	// The cases where this happens includes:
	// - The jInX node places an order in it's in-memory orderbook during the
	//   CheckTx flow.
	// A best effort placed order may not have been placed on other jInX
	// nodes including other jInX validator nodes and may still be excluded in
	// future order matches.
	OrderPlaceV1_ORDER_PLACEMENT_STATUS_BEST_EFFORT_OPENED OrderPlaceV1_OrderPlacementStatus = 1
	// An opened order is one that is confirmed to be placed on all jInX nodes
	// (discounting dishonest jInX nodes) and will be included in any future
	// order matches.
	// This status is used internally by the indexer and will not be sent
	// out by protocol.
	OrderPlaceV1_ORDER_PLACEMENT_STATUS_OPENED OrderPlaceV1_OrderPlacementStatus = 2
)

var OrderPlaceV1_OrderPlacementStatus_name = map[int32]string{
	0: "ORDER_PLACEMENT_STATUS_UNSPECIFIED",
	1: "ORDER_PLACEMENT_STATUS_BEST_EFFORT_OPENED",
	2: "ORDER_PLACEMENT_STATUS_OPENED",
}

var OrderPlaceV1_OrderPlacementStatus_value = map[string]int32{
	"ORDER_PLACEMENT_STATUS_UNSPECIFIED":        0,
	"ORDER_PLACEMENT_STATUS_BEST_EFFORT_OPENED": 1,
	"ORDER_PLACEMENT_STATUS_OPENED":             2,
}

func (x OrderPlaceV1_OrderPlacementStatus) String() string {
	return proto.EnumName(OrderPlaceV1_OrderPlacementStatus_name, int32(x))
}

func (OrderPlaceV1_OrderPlacementStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{0, 0}
}

// OrderRemovalStatus is an enum for the resulting status after an order is
// removed.
type OrderRemoveV1_OrderRemovalStatus int32

const (
	// Default value, this is invalid and unused.
	OrderRemoveV1_ORDER_REMOVAL_STATUS_UNSPECIFIED OrderRemoveV1_OrderRemovalStatus = 0
	// A best effort canceled order is one that has only been confirmed to be
	// removed on the jInX node sending the off-chain update message.
	// The cases where this happens includes:
	// - the order was removed due to the jInX node receiving a CancelOrder
	//   transaction for the order.
	// - the order was removed due to being undercollateralized during
	//   optimistic matching.
	// A best effort canceled order may not have been removed on other jInX
	// nodes including other jInX validator nodes and may still be included in
	// future order matches.
	OrderRemoveV1_ORDER_REMOVAL_STATUS_BEST_EFFORT_CANCELED OrderRemoveV1_OrderRemovalStatus = 1
	// A canceled order is one that is confirmed to be removed on all jInX nodes
	// (discounting dishonest jInX nodes) and will not be included in any future
	// order matches.
	// The cases where this happens includes:
	// - the order is expired.
	OrderRemoveV1_ORDER_REMOVAL_STATUS_CANCELED OrderRemoveV1_OrderRemovalStatus = 2
	// An order was fully-filled. Only sent by the Indexer for stateful orders.
	OrderRemoveV1_ORDER_REMOVAL_STATUS_FILLED OrderRemoveV1_OrderRemovalStatus = 3
)

var OrderRemoveV1_OrderRemovalStatus_name = map[int32]string{
	0: "ORDER_REMOVAL_STATUS_UNSPECIFIED",
	1: "ORDER_REMOVAL_STATUS_BEST_EFFORT_CANCELED",
	2: "ORDER_REMOVAL_STATUS_CANCELED",
	3: "ORDER_REMOVAL_STATUS_FILLED",
}

var OrderRemoveV1_OrderRemovalStatus_value = map[string]int32{
	"ORDER_REMOVAL_STATUS_UNSPECIFIED":          0,
	"ORDER_REMOVAL_STATUS_BEST_EFFORT_CANCELED": 1,
	"ORDER_REMOVAL_STATUS_CANCELED":             2,
	"ORDER_REMOVAL_STATUS_FILLED":               3,
}

func (x OrderRemoveV1_OrderRemovalStatus) String() string {
	return proto.EnumName(OrderRemoveV1_OrderRemovalStatus_name, int32(x))
}

func (OrderRemoveV1_OrderRemovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{1, 0}
}

// OrderPlace messages contain the order placed/replaced.
type OrderPlaceV1 struct {
	Order           *v1.IndexerOrder                  `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	PlacementStatus OrderPlaceV1_OrderPlacementStatus `protobuf:"varint,2,opt,name=placement_status,json=placementStatus,proto3,enum=jinxprotocol.indexer.off_chain_updates.OrderPlaceV1_OrderPlacementStatus" json:"placement_status,omitempty"`
}

func (m *OrderPlaceV1) Reset()         { *m = OrderPlaceV1{} }
func (m *OrderPlaceV1) String() string { return proto.CompactTextString(m) }
func (*OrderPlaceV1) ProtoMessage()    {}
func (*OrderPlaceV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{0}
}
func (m *OrderPlaceV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderPlaceV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderPlaceV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderPlaceV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPlaceV1.Merge(m, src)
}
func (m *OrderPlaceV1) XXX_Size() int {
	return m.Size()
}
func (m *OrderPlaceV1) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPlaceV1.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPlaceV1 proto.InternalMessageInfo

func (m *OrderPlaceV1) GetOrder() *v1.IndexerOrder {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderPlaceV1) GetPlacementStatus() OrderPlaceV1_OrderPlacementStatus {
	if m != nil {
		return m.PlacementStatus
	}
	return OrderPlaceV1_ORDER_PLACEMENT_STATUS_UNSPECIFIED
}

// OrderRemove messages contain the id of the order removed, the reason for the
// removal and the resulting status from the removal.
type OrderRemoveV1 struct {
	RemovedOrderId *v1.IndexerOrderId               `protobuf:"bytes,1,opt,name=removed_order_id,json=removedOrderId,proto3" json:"removed_order_id,omitempty"`
	Reason         shared.OrderRemovalReason        `protobuf:"varint,2,opt,name=reason,proto3,enum=jinxprotocol.indexer.shared.OrderRemovalReason" json:"reason,omitempty"`
	RemovalStatus  OrderRemoveV1_OrderRemovalStatus `protobuf:"varint,3,opt,name=removal_status,json=removalStatus,proto3,enum=jinxprotocol.indexer.off_chain_updates.OrderRemoveV1_OrderRemovalStatus" json:"removal_status,omitempty"`
}

func (m *OrderRemoveV1) Reset()         { *m = OrderRemoveV1{} }
func (m *OrderRemoveV1) String() string { return proto.CompactTextString(m) }
func (*OrderRemoveV1) ProtoMessage()    {}
func (*OrderRemoveV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{1}
}
func (m *OrderRemoveV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderRemoveV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderRemoveV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderRemoveV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRemoveV1.Merge(m, src)
}
func (m *OrderRemoveV1) XXX_Size() int {
	return m.Size()
}
func (m *OrderRemoveV1) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRemoveV1.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRemoveV1 proto.InternalMessageInfo

func (m *OrderRemoveV1) GetRemovedOrderId() *v1.IndexerOrderId {
	if m != nil {
		return m.RemovedOrderId
	}
	return nil
}

func (m *OrderRemoveV1) GetReason() shared.OrderRemovalReason {
	if m != nil {
		return m.Reason
	}
	return shared.OrderRemovalReason_ORDER_REMOVAL_REASON_UNSPECIFIED
}

func (m *OrderRemoveV1) GetRemovalStatus() OrderRemoveV1_OrderRemovalStatus {
	if m != nil {
		return m.RemovalStatus
	}
	return OrderRemoveV1_ORDER_REMOVAL_STATUS_UNSPECIFIED
}

// OrderUpdate messages contain the id of the order being updated, and the
// updated total filled quantums of the order.
type OrderUpdateV1 struct {
	OrderId             *v1.IndexerOrderId `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TotalFilledQuantums uint64             `protobuf:"varint,2,opt,name=total_filled_quantums,json=totalFilledQuantums,proto3" json:"total_filled_quantums,omitempty"`
}

func (m *OrderUpdateV1) Reset()         { *m = OrderUpdateV1{} }
func (m *OrderUpdateV1) String() string { return proto.CompactTextString(m) }
func (*OrderUpdateV1) ProtoMessage()    {}
func (*OrderUpdateV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{2}
}
func (m *OrderUpdateV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderUpdateV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderUpdateV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderUpdateV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderUpdateV1.Merge(m, src)
}
func (m *OrderUpdateV1) XXX_Size() int {
	return m.Size()
}
func (m *OrderUpdateV1) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderUpdateV1.DiscardUnknown(m)
}

var xxx_messageInfo_OrderUpdateV1 proto.InternalMessageInfo

func (m *OrderUpdateV1) GetOrderId() *v1.IndexerOrderId {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func (m *OrderUpdateV1) GetTotalFilledQuantums() uint64 {
	if m != nil {
		return m.TotalFilledQuantums
	}
	return 0
}

// An OffChainUpdate message is the message type which will be sent on Kafka to
// the Indexer.
type OffChainUpdateV1 struct {
	// Contains one of an OrderPlaceV1, OrderRemoveV1, and OrderUpdateV1 message.
	//
	// Types that are valid to be assigned to UpdateMessage:
	//	*OffChainUpdateV1_OrderPlace
	//	*OffChainUpdateV1_OrderRemove
	//	*OffChainUpdateV1_OrderUpdate
	UpdateMessage isOffChainUpdateV1_UpdateMessage `protobuf_oneof:"update_message"`
}

func (m *OffChainUpdateV1) Reset()         { *m = OffChainUpdateV1{} }
func (m *OffChainUpdateV1) String() string { return proto.CompactTextString(m) }
func (*OffChainUpdateV1) ProtoMessage()    {}
func (*OffChainUpdateV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4bf75e7277ef66, []int{3}
}
func (m *OffChainUpdateV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OffChainUpdateV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OffChainUpdateV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OffChainUpdateV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OffChainUpdateV1.Merge(m, src)
}
func (m *OffChainUpdateV1) XXX_Size() int {
	return m.Size()
}
func (m *OffChainUpdateV1) XXX_DiscardUnknown() {
	xxx_messageInfo_OffChainUpdateV1.DiscardUnknown(m)
}

var xxx_messageInfo_OffChainUpdateV1 proto.InternalMessageInfo

type isOffChainUpdateV1_UpdateMessage interface {
	isOffChainUpdateV1_UpdateMessage()
	MarshalTo([]byte) (int, error)
	Size() int
}

type OffChainUpdateV1_OrderPlace struct {
	OrderPlace *OrderPlaceV1 `protobuf:"bytes,1,opt,name=order_place,json=orderPlace,proto3,oneof" json:"order_place,omitempty"`
}
type OffChainUpdateV1_OrderRemove struct {
	OrderRemove *OrderRemoveV1 `protobuf:"bytes,2,opt,name=order_remove,json=orderRemove,proto3,oneof" json:"order_remove,omitempty"`
}
type OffChainUpdateV1_OrderUpdate struct {
	OrderUpdate *OrderUpdateV1 `protobuf:"bytes,3,opt,name=order_update,json=orderUpdate,proto3,oneof" json:"order_update,omitempty"`
}

func (*OffChainUpdateV1_OrderPlace) isOffChainUpdateV1_UpdateMessage()  {}
func (*OffChainUpdateV1_OrderRemove) isOffChainUpdateV1_UpdateMessage() {}
func (*OffChainUpdateV1_OrderUpdate) isOffChainUpdateV1_UpdateMessage() {}

func (m *OffChainUpdateV1) GetUpdateMessage() isOffChainUpdateV1_UpdateMessage {
	if m != nil {
		return m.UpdateMessage
	}
	return nil
}

func (m *OffChainUpdateV1) GetOrderPlace() *OrderPlaceV1 {
	if x, ok := m.GetUpdateMessage().(*OffChainUpdateV1_OrderPlace); ok {
		return x.OrderPlace
	}
	return nil
}

func (m *OffChainUpdateV1) GetOrderRemove() *OrderRemoveV1 {
	if x, ok := m.GetUpdateMessage().(*OffChainUpdateV1_OrderRemove); ok {
		return x.OrderRemove
	}
	return nil
}

func (m *OffChainUpdateV1) GetOrderUpdate() *OrderUpdateV1 {
	if x, ok := m.GetUpdateMessage().(*OffChainUpdateV1_OrderUpdate); ok {
		return x.OrderUpdate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OffChainUpdateV1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OffChainUpdateV1_OrderPlace)(nil),
		(*OffChainUpdateV1_OrderRemove)(nil),
		(*OffChainUpdateV1_OrderUpdate)(nil),
	}
}

func init() {
	proto.RegisterEnum("jinxprotocol.indexer.off_chain_updates.OrderPlaceV1_OrderPlacementStatus", OrderPlaceV1_OrderPlacementStatus_name, OrderPlaceV1_OrderPlacementStatus_value)
	proto.RegisterEnum("jinxprotocol.indexer.off_chain_updates.OrderRemoveV1_OrderRemovalStatus", OrderRemoveV1_OrderRemovalStatus_name, OrderRemoveV1_OrderRemovalStatus_value)
	proto.RegisterType((*OrderPlaceV1)(nil), "jinxprotocol.indexer.off_chain_updates.OrderPlaceV1")
	proto.RegisterType((*OrderRemoveV1)(nil), "jinxprotocol.indexer.off_chain_updates.OrderRemoveV1")
	proto.RegisterType((*OrderUpdateV1)(nil), "jinxprotocol.indexer.off_chain_updates.OrderUpdateV1")
	proto.RegisterType((*OffChainUpdateV1)(nil), "jinxprotocol.indexer.off_chain_updates.OffChainUpdateV1")
}

func init() {
	proto.RegisterFile("jinxprotocol/indexer/off_chain_updates/off_chain_updates.proto", fileDescriptor_6e4bf75e7277ef66)
}

var fileDescriptor_6e4bf75e7277ef66 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x6e, 0xd3, 0x3c,
	0x18, 0x4d, 0xba, 0xff, 0x1f, 0xc8, 0xdb, 0x4a, 0x64, 0x40, 0x9a, 0x86, 0x08, 0x23, 0x42, 0xd3,
	0x10, 0x5a, 0xb2, 0x96, 0x71, 0x8b, 0xd4, 0xa5, 0x29, 0x8b, 0xe8, 0x9a, 0xe2, 0x76, 0x43, 0x9a,
	0x84, 0xac, 0xac, 0x71, 0xb7, 0xa0, 0x34, 0x2e, 0x49, 0x5a, 0xed, 0x31, 0xf6, 0x22, 0x5c, 0xf2,
	0x0e, 0x5c, 0xee, 0x06, 0x89, 0x1b, 0x24, 0xb4, 0xbe, 0x08, 0x8a, 0xed, 0x66, 0xed, 0x9a, 0x69,
	0x74, 0x5c, 0x7e, 0x5f, 0xce, 0x77, 0x7c, 0x7c, 0xce, 0xe7, 0x16, 0xbc, 0xfd, 0xec, 0x87, 0x67,
	0xfd, 0x88, 0x26, 0xb4, 0x43, 0x03, 0xc3, 0x0f, 0x3d, 0x72, 0x46, 0x22, 0x83, 0x76, 0xbb, 0xb8,
	0x73, 0xea, 0xfa, 0x21, 0x1e, 0xf4, 0x3d, 0x37, 0x21, 0xf1, 0x6c, 0x47, 0x67, 0x43, 0x70, 0x63,
	0x72, 0x5e, 0x17, 0xf3, 0xfa, 0x0c, 0x7a, 0x6d, 0x3b, 0xf7, 0x9c, 0xf8, 0xd4, 0x8d, 0x88, 0x67,
	0x44, 0xa4, 0x47, 0x87, 0x6e, 0x80, 0x23, 0xe2, 0xc6, 0x34, 0xe4, 0xcc, 0x6b, 0xaf, 0x72, 0x27,
	0xb2, 0xc6, 0xb0, 0x64, 0x74, 0x02, 0x7a, 0xcc, 0xc1, 0xda, 0xaf, 0x02, 0x58, 0x76, 0x22, 0x8f,
	0x44, 0xcd, 0xc0, 0xed, 0x90, 0xc3, 0x12, 0xac, 0x82, 0xff, 0x69, 0x5a, 0xaf, 0xca, 0xeb, 0xf2,
	0xe6, 0x52, 0x59, 0xd7, 0x73, 0x75, 0x66, 0x8d, 0x61, 0x49, 0xb7, 0x79, 0x8f, 0xb1, 0x20, 0x3e,
	0x0c, 0x13, 0xa0, 0xf4, 0x53, 0xc2, 0x1e, 0x09, 0x13, 0x1c, 0x27, 0x6e, 0x32, 0x88, 0x57, 0x0b,
	0xeb, 0xf2, 0x66, 0xb1, 0x6c, 0xeb, 0x7f, 0x77, 0x71, 0x7d, 0x52, 0xd5, 0x44, 0x91, 0x32, 0xb6,
	0x18, 0x21, 0x7a, 0xd0, 0x9f, 0x6e, 0x68, 0xe7, 0x32, 0x78, 0x94, 0x87, 0x84, 0x1b, 0x40, 0x73,
	0x50, 0xd5, 0x42, 0xb8, 0x59, 0xaf, 0x98, 0xd6, 0xbe, 0xd5, 0x68, 0xe3, 0x56, 0xbb, 0xd2, 0x3e,
	0x68, 0xe1, 0x83, 0x46, 0xab, 0x69, 0x99, 0x76, 0xcd, 0xb6, 0xaa, 0x8a, 0x04, 0xb7, 0xc0, 0xcb,
	0x1b, 0x70, 0xbb, 0x56, 0xab, 0x8d, 0xad, 0x5a, 0xcd, 0x41, 0x6d, 0xec, 0x34, 0xad, 0x86, 0x55,
	0x55, 0x64, 0xf8, 0x1c, 0x3c, 0xbd, 0x01, 0x2e, 0x20, 0x05, 0xed, 0xc7, 0x02, 0x58, 0xe1, 0xce,
	0xa4, 0x51, 0xa5, 0x06, 0x1f, 0x01, 0x85, 0xc5, 0x46, 0x3c, 0xcc, 0xbc, 0xc2, 0xbe, 0x27, 0xbc,
	0xde, 0x9e, 0xcf, 0x6b, 0xdb, 0x43, 0x45, 0xc1, 0x24, 0x6a, 0xf8, 0x0e, 0x2c, 0xf2, 0x55, 0x10,
	0x66, 0x1b, 0xf9, 0x8c, 0x7c, 0x7b, 0xf4, 0x2b, 0x5d, 0x6e, 0x80, 0xd8, 0x18, 0x12, 0xe3, 0x90,
	0x82, 0xe2, 0x78, 0xb7, 0x44, 0x7a, 0x0b, 0x8c, 0x70, 0x6f, 0xae, 0xf4, 0xc6, 0x77, 0x9e, 0x3a,
	0x49, 0x84, 0xb7, 0x12, 0x4d, 0x96, 0xda, 0x57, 0x19, 0xc0, 0x59, 0x14, 0x7c, 0x01, 0xd6, 0xb9,
	0xc3, 0xc8, 0xda, 0x77, 0x0e, 0x2b, 0xf5, 0x5b, 0x62, 0xbb, 0x86, 0x9a, 0x0c, 0xcd, 0xac, 0x34,
	0x4c, 0xab, 0x3e, 0x1d, 0xdb, 0x35, 0x78, 0x06, 0x29, 0xc0, 0x67, 0xe0, 0x49, 0x2e, 0xa4, 0x66,
	0xd7, 0x53, 0xc0, 0x42, 0xba, 0x6a, 0x3c, 0xd7, 0x03, 0x76, 0xe1, 0xc3, 0x12, 0x7c, 0x0f, 0xee,
	0xff, 0x73, 0x9e, 0xf7, 0xa8, 0x08, 0xb2, 0x0c, 0x1e, 0x27, 0x34, 0x71, 0x03, 0xdc, 0xf5, 0x83,
	0x80, 0x78, 0xf8, 0xcb, 0xc0, 0x0d, 0x93, 0x41, 0x8f, 0x3f, 0xa2, 0xff, 0xd0, 0x43, 0xf6, 0xb1,
	0xc6, 0xbe, 0x7d, 0x10, 0x9f, 0xb4, 0x6f, 0x05, 0xa0, 0x38, 0xdd, 0xae, 0x99, 0xe6, 0x90, 0xa9,
	0xfa, 0x08, 0x96, 0xb8, 0x2a, 0xf6, 0x56, 0x84, 0xb0, 0x9d, 0xbb, 0xbc, 0xc1, 0x3d, 0x09, 0x01,
	0x9a, 0xd5, 0xf0, 0x08, 0x2c, 0x73, 0x62, 0xbe, 0x82, 0x4c, 0xd8, 0x52, 0xf9, 0xcd, 0x9d, 0xf6,
	0x63, 0x4f, 0x42, 0x5c, 0x25, 0x6f, 0x5c, 0x71, 0x73, 0x34, 0xdb, 0xbd, 0x79, 0xb9, 0xc7, 0x0e,
	0x64, 0xdc, 0xbc, 0xb1, 0xab, 0x80, 0x22, 0xc7, 0xe1, 0x1e, 0x89, 0x63, 0xf7, 0x84, 0xec, 0x7e,
	0xfa, 0x7e, 0xa9, 0xca, 0x17, 0x97, 0xaa, 0xfc, 0xfb, 0x52, 0x95, 0xcf, 0x47, 0xaa, 0x74, 0x31,
	0x52, 0xa5, 0x9f, 0x23, 0x55, 0x3a, 0x32, 0x4f, 0xfc, 0xe4, 0x74, 0x70, 0xac, 0x77, 0x68, 0xcf,
	0x98, 0xfa, 0x51, 0x1d, 0xee, 0x6c, 0xb1, 0x23, 0x8d, 0xdb, 0xff, 0x00, 0x8e, 0x17, 0x19, 0xe6,
	0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x84, 0x8e, 0xaf, 0x31, 0x06, 0x00, 0x00,
}

func (m *OrderPlaceV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderPlaceV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderPlaceV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlacementStatus != 0 {
		i = encodeVarintOffChainUpdates(dAtA, i, uint64(m.PlacementStatus))
		i--
		dAtA[i] = 0x10
	}
	if m.Order != nil {
		{
			size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrderRemoveV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderRemoveV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderRemoveV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RemovalStatus != 0 {
		i = encodeVarintOffChainUpdates(dAtA, i, uint64(m.RemovalStatus))
		i--
		dAtA[i] = 0x18
	}
	if m.Reason != 0 {
		i = encodeVarintOffChainUpdates(dAtA, i, uint64(m.Reason))
		i--
		dAtA[i] = 0x10
	}
	if m.RemovedOrderId != nil {
		{
			size, err := m.RemovedOrderId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrderUpdateV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderUpdateV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderUpdateV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalFilledQuantums != 0 {
		i = encodeVarintOffChainUpdates(dAtA, i, uint64(m.TotalFilledQuantums))
		i--
		dAtA[i] = 0x10
	}
	if m.OrderId != nil {
		{
			size, err := m.OrderId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OffChainUpdateV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OffChainUpdateV1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OffChainUpdateV1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateMessage != nil {
		{
			size := m.UpdateMessage.Size()
			i -= size
			if _, err := m.UpdateMessage.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *OffChainUpdateV1_OrderPlace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OffChainUpdateV1_OrderPlace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderPlace != nil {
		{
			size, err := m.OrderPlace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *OffChainUpdateV1_OrderRemove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OffChainUpdateV1_OrderRemove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderRemove != nil {
		{
			size, err := m.OrderRemove.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *OffChainUpdateV1_OrderUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OffChainUpdateV1_OrderUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderUpdate != nil {
		{
			size, err := m.OrderUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOffChainUpdates(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintOffChainUpdates(dAtA []byte, offset int, v uint64) int {
	offset -= sovOffChainUpdates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderPlaceV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Order != nil {
		l = m.Order.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	if m.PlacementStatus != 0 {
		n += 1 + sovOffChainUpdates(uint64(m.PlacementStatus))
	}
	return n
}

func (m *OrderRemoveV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemovedOrderId != nil {
		l = m.RemovedOrderId.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	if m.Reason != 0 {
		n += 1 + sovOffChainUpdates(uint64(m.Reason))
	}
	if m.RemovalStatus != 0 {
		n += 1 + sovOffChainUpdates(uint64(m.RemovalStatus))
	}
	return n
}

func (m *OrderUpdateV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderId != nil {
		l = m.OrderId.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	if m.TotalFilledQuantums != 0 {
		n += 1 + sovOffChainUpdates(uint64(m.TotalFilledQuantums))
	}
	return n
}

func (m *OffChainUpdateV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateMessage != nil {
		n += m.UpdateMessage.Size()
	}
	return n
}

func (m *OffChainUpdateV1_OrderPlace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderPlace != nil {
		l = m.OrderPlace.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	return n
}
func (m *OffChainUpdateV1_OrderRemove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderRemove != nil {
		l = m.OrderRemove.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	return n
}
func (m *OffChainUpdateV1_OrderUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderUpdate != nil {
		l = m.OrderUpdate.Size()
		n += 1 + l + sovOffChainUpdates(uint64(l))
	}
	return n
}

func sovOffChainUpdates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOffChainUpdates(x uint64) (n int) {
	return sovOffChainUpdates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderPlaceV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffChainUpdates
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
			return fmt.Errorf("proto: OrderPlaceV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderPlaceV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Order == nil {
				m.Order = &v1.IndexerOrder{}
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlacementStatus", wireType)
			}
			m.PlacementStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlacementStatus |= OrderPlaceV1_OrderPlacementStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOffChainUpdates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffChainUpdates
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
func (m *OrderRemoveV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffChainUpdates
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
			return fmt.Errorf("proto: OrderRemoveV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderRemoveV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedOrderId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemovedOrderId == nil {
				m.RemovedOrderId = &v1.IndexerOrderId{}
			}
			if err := m.RemovedOrderId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			m.Reason = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reason |= shared.OrderRemovalReason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovalStatus", wireType)
			}
			m.RemovalStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemovalStatus |= OrderRemoveV1_OrderRemovalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOffChainUpdates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffChainUpdates
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
func (m *OrderUpdateV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffChainUpdates
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
			return fmt.Errorf("proto: OrderUpdateV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderUpdateV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OrderId == nil {
				m.OrderId = &v1.IndexerOrderId{}
			}
			if err := m.OrderId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFilledQuantums", wireType)
			}
			m.TotalFilledQuantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalFilledQuantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOffChainUpdates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffChainUpdates
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
func (m *OffChainUpdateV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOffChainUpdates
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
			return fmt.Errorf("proto: OffChainUpdateV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OffChainUpdateV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderPlace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OrderPlaceV1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpdateMessage = &OffChainUpdateV1_OrderPlace{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderRemove", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OrderRemoveV1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpdateMessage = &OffChainUpdateV1_OrderRemove{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOffChainUpdates
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
				return ErrInvalidLengthOffChainUpdates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOffChainUpdates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OrderUpdateV1{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.UpdateMessage = &OffChainUpdateV1_OrderUpdate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOffChainUpdates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOffChainUpdates
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
func skipOffChainUpdates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOffChainUpdates
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
					return 0, ErrIntOverflowOffChainUpdates
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
					return 0, ErrIntOverflowOffChainUpdates
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
				return 0, ErrInvalidLengthOffChainUpdates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOffChainUpdates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOffChainUpdates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOffChainUpdates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOffChainUpdates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOffChainUpdates = fmt.Errorf("proto: unexpected end of group")
)
