// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/vest/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgDeleteVestEntry is the Msg/DeleteVestEntry request type.
type MsgDeleteVestEntry struct {
	// authority is the address that controls the module.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// The vester account of the vest entry to delete.
	VesterAccount string `protobuf:"bytes,2,opt,name=vester_account,json=vesterAccount,proto3" json:"vester_account,omitempty"`
}

func (m *MsgDeleteVestEntry) Reset()         { *m = MsgDeleteVestEntry{} }
func (m *MsgDeleteVestEntry) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteVestEntry) ProtoMessage()    {}
func (*MsgDeleteVestEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc12a3a68223b96, []int{0}
}
func (m *MsgDeleteVestEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteVestEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteVestEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteVestEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteVestEntry.Merge(m, src)
}
func (m *MsgDeleteVestEntry) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteVestEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteVestEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteVestEntry proto.InternalMessageInfo

func (m *MsgDeleteVestEntry) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgDeleteVestEntry) GetVesterAccount() string {
	if m != nil {
		return m.VesterAccount
	}
	return ""
}

// MsgDeleteVestEntryResponse is the Msg/DeleteVestEntry response type.
type MsgDeleteVestEntryResponse struct {
}

func (m *MsgDeleteVestEntryResponse) Reset()         { *m = MsgDeleteVestEntryResponse{} }
func (m *MsgDeleteVestEntryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteVestEntryResponse) ProtoMessage()    {}
func (*MsgDeleteVestEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc12a3a68223b96, []int{1}
}
func (m *MsgDeleteVestEntryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteVestEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteVestEntryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteVestEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteVestEntryResponse.Merge(m, src)
}
func (m *MsgDeleteVestEntryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteVestEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteVestEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteVestEntryResponse proto.InternalMessageInfo

// MsgSetVestEntry is the Msg/SetVestEntry request type.
type MsgSetVestEntry struct {
	// authority is the address that controls the module.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// The vest entry to set.
	Entry VestEntry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry"`
}

func (m *MsgSetVestEntry) Reset()         { *m = MsgSetVestEntry{} }
func (m *MsgSetVestEntry) String() string { return proto.CompactTextString(m) }
func (*MsgSetVestEntry) ProtoMessage()    {}
func (*MsgSetVestEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc12a3a68223b96, []int{2}
}
func (m *MsgSetVestEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetVestEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetVestEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetVestEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetVestEntry.Merge(m, src)
}
func (m *MsgSetVestEntry) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetVestEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetVestEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetVestEntry proto.InternalMessageInfo

func (m *MsgSetVestEntry) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetVestEntry) GetEntry() VestEntry {
	if m != nil {
		return m.Entry
	}
	return VestEntry{}
}

// MsgSetVestEntryResponse is the Msg/SetVestEntry response type.
type MsgSetVestEntryResponse struct {
}

func (m *MsgSetVestEntryResponse) Reset()         { *m = MsgSetVestEntryResponse{} }
func (m *MsgSetVestEntryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetVestEntryResponse) ProtoMessage()    {}
func (*MsgSetVestEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc12a3a68223b96, []int{3}
}
func (m *MsgSetVestEntryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetVestEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetVestEntryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetVestEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetVestEntryResponse.Merge(m, src)
}
func (m *MsgSetVestEntryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetVestEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetVestEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetVestEntryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDeleteVestEntry)(nil), "jinxprotocol.vest.MsgDeleteVestEntry")
	proto.RegisterType((*MsgDeleteVestEntryResponse)(nil), "jinxprotocol.vest.MsgDeleteVestEntryResponse")
	proto.RegisterType((*MsgSetVestEntry)(nil), "jinxprotocol.vest.MsgSetVestEntry")
	proto.RegisterType((*MsgSetVestEntryResponse)(nil), "jinxprotocol.vest.MsgSetVestEntryResponse")
}

func init() { proto.RegisterFile("jinxprotocol/vest/tx.proto", fileDescriptor_2dc12a3a68223b96) }

var fileDescriptor_2dc12a3a68223b96 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xca, 0xcc, 0xab,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4b, 0x2d, 0x2e, 0xd1, 0x2f, 0xa9,
	0xd0, 0x03, 0x0b, 0x08, 0x09, 0x22, 0xcb, 0xe9, 0x81, 0xe4, 0xa4, 0x24, 0x93, 0xf3, 0x8b, 0x73,
	0xf3, 0x8b, 0xe3, 0xc1, 0xa2, 0xfa, 0x10, 0x0e, 0x44, 0xb5, 0x94, 0x38, 0x84, 0xa7, 0x9f, 0x5b,
	0x9c, 0xae, 0x5f, 0x66, 0x08, 0xa2, 0xa0, 0x12, 0x4a, 0x98, 0x56, 0x80, 0x88, 0xf8, 0xd4, 0xbc,
	0x92, 0xa2, 0x4a, 0xa8, 0x1a, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x88, 0xa1, 0x20, 0x16, 0x44, 0x54,
	0xa9, 0x99, 0x91, 0x4b, 0xc8, 0xb7, 0x38, 0xdd, 0x25, 0x35, 0x27, 0xb5, 0x24, 0x35, 0x2c, 0xb5,
	0xb8, 0xc4, 0x15, 0xa4, 0x45, 0xc8, 0x8c, 0x8b, 0x33, 0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x28, 0xb3,
	0xa4, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xe2, 0xd2, 0x16, 0x5d, 0x11, 0xa8, 0x73,
	0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x83, 0x10, 0x4a,
	0x85, 0x54, 0xb9, 0xf8, 0x40, 0x16, 0xa7, 0x16, 0xc5, 0x27, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95,
	0x48, 0x30, 0x81, 0x34, 0x07, 0xf1, 0x42, 0x44, 0x1d, 0x21, 0x82, 0x56, 0x7c, 0x4d, 0xcf, 0x37,
	0x68, 0x21, 0xb4, 0x29, 0xc9, 0x70, 0x49, 0x61, 0x3a, 0x22, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x69, 0x32, 0x23, 0x17, 0xbf, 0x6f, 0x71, 0x7a, 0x70, 0x6a, 0x09, 0xe5, 0x0e, 0xb4,
	0xe0, 0x62, 0x05, 0x07, 0x0a, 0xd8, 0x5d, 0xdc, 0x46, 0x32, 0x7a, 0x18, 0x11, 0xa0, 0x07, 0xb7,
	0xc4, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x88, 0x06, 0x0c, 0x37, 0x4b, 0x72, 0x89, 0xa3,
	0x39, 0x0a, 0xe6, 0x60, 0xa3, 0x73, 0x8c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x71, 0x5c, 0x3c,
	0x28, 0x8e, 0x56, 0xc2, 0x62, 0x1b, 0x9a, 0x19, 0x52, 0x5a, 0x84, 0xd5, 0xc0, 0xec, 0x11, 0x4a,
	0xe7, 0xe2, 0x47, 0x8f, 0x38, 0x55, 0xec, 0xda, 0xd1, 0x94, 0x49, 0xe9, 0x12, 0xa5, 0x0c, 0x66,
	0x91, 0x53, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa5, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa3, 0x26, 0x42, 0x13, 0xdd, 0xe4, 0x8c,
	0xc4, 0xcc, 0x3c, 0x7d, 0xb8, 0x48, 0x05, 0x34, 0xed, 0x57, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81,
	0x85, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x62, 0x14, 0x02, 0x1d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SetVestEntry sets a VestEntry in state.
	SetVestEntry(ctx context.Context, in *MsgSetVestEntry, opts ...grpc.CallOption) (*MsgSetVestEntryResponse, error)
	// DeleteVestEntry deletes a VestEntry from state.
	DeleteVestEntry(ctx context.Context, in *MsgDeleteVestEntry, opts ...grpc.CallOption) (*MsgDeleteVestEntryResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetVestEntry(ctx context.Context, in *MsgSetVestEntry, opts ...grpc.CallOption) (*MsgSetVestEntryResponse, error) {
	out := new(MsgSetVestEntryResponse)
	err := c.cc.Invoke(ctx, "/jinxprotocol.vest.Msg/SetVestEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteVestEntry(ctx context.Context, in *MsgDeleteVestEntry, opts ...grpc.CallOption) (*MsgDeleteVestEntryResponse, error) {
	out := new(MsgDeleteVestEntryResponse)
	err := c.cc.Invoke(ctx, "/jinxprotocol.vest.Msg/DeleteVestEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetVestEntry sets a VestEntry in state.
	SetVestEntry(context.Context, *MsgSetVestEntry) (*MsgSetVestEntryResponse, error)
	// DeleteVestEntry deletes a VestEntry from state.
	DeleteVestEntry(context.Context, *MsgDeleteVestEntry) (*MsgDeleteVestEntryResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetVestEntry(ctx context.Context, req *MsgSetVestEntry) (*MsgSetVestEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVestEntry not implemented")
}
func (*UnimplementedMsgServer) DeleteVestEntry(ctx context.Context, req *MsgDeleteVestEntry) (*MsgDeleteVestEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVestEntry not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetVestEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetVestEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetVestEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jinxprotocol.vest.Msg/SetVestEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetVestEntry(ctx, req.(*MsgSetVestEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteVestEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteVestEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteVestEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jinxprotocol.vest.Msg/DeleteVestEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteVestEntry(ctx, req.(*MsgDeleteVestEntry))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jinxprotocol.vest.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetVestEntry",
			Handler:    _Msg_SetVestEntry_Handler,
		},
		{
			MethodName: "DeleteVestEntry",
			Handler:    _Msg_DeleteVestEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jinxprotocol/vest/tx.proto",
}

func (m *MsgDeleteVestEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteVestEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteVestEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VesterAccount) > 0 {
		i -= len(m.VesterAccount)
		copy(dAtA[i:], m.VesterAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VesterAccount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteVestEntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteVestEntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteVestEntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetVestEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetVestEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetVestEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Entry.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetVestEntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetVestEntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetVestEntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDeleteVestEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VesterAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteVestEntryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetVestEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Entry.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetVestEntryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeleteVestEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteVestEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteVestEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VesterAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VesterAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteVestEntryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteVestEntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteVestEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetVestEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetVestEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetVestEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetVestEntryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetVestEntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetVestEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
